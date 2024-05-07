package emby

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	pb "github.com/synctv-org/vendors/api/emby"
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/vendors/emby"
)

type EmbyService struct {
	pb.UnimplementedEmbyServer
}

func NewEmbyService(c *conf.EmbyConfig) *EmbyService {
	return &EmbyService{}
}

func (a *EmbyService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	cli := emby.NewClient(req.Host, emby.WithContext(ctx))
	r, err := cli.GetApiKey(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		Token:    r.AccessToken,
		UserId:   r.UserID,
		ServerId: r.ServerID,
	}, nil
}

func (a *EmbyService) Me(ctx context.Context, req *pb.MeReq) (*pb.MeResp, error) {
	cli := emby.NewClient(req.Host, emby.WithContext(ctx), emby.WithKey(req.Token), emby.WithUserID(req.UserId))
	r, err := cli.Me()
	if err != nil {
		return nil, err
	}
	return &pb.MeResp{
		Id:       r.ID,
		Name:     r.Name,
		ServerId: r.ServerID,
	}, nil
}

func mediaStreamInfo2pb(msi []emby.MediaStreams) []*pb.MediaStreamInfo {
	var pbmsi []*pb.MediaStreamInfo = make([]*pb.MediaStreamInfo, len(msi))
	for i, msi := range msi {
		pbmsi[i] = &pb.MediaStreamInfo{
			Codec:           msi.Codec,
			Language:        msi.Language,
			Type:            msi.Type,
			Title:           msi.Title,
			DisplayTitle:    msi.DisplayTitle,
			DisplayLanguage: msi.DisplayLanguage,
			IsDefault:       msi.IsDefault,
			Index:           msi.Index,
			Protocol:        msi.Protocol,
		}
	}
	return pbmsi
}

func mediaSources2pb(ms []emby.MediaSources) []*pb.MediaSourceInfo {
	var pbms []*pb.MediaSourceInfo = make([]*pb.MediaSourceInfo, len(ms))
	for i, msi := range ms {
		pbms[i] = &pb.MediaSourceInfo{
			Id:                         msi.ID,
			Name:                       msi.Name,
			Path:                       msi.Path,
			Protocol:                   msi.Protocol,
			Container:                  msi.Container,
			DefaultSubtitleStreamIndex: msi.DefaultSubtitleStreamIndex,
			DefaultAudioStreamIndex:    msi.DefaultAudioStreamIndex,
			MediaStreamInfo:            mediaStreamInfo2pb(msi.MediaStreams),
			DirectPlayUrl:              msi.DirectStreamURL,
			TranscodingUrl:             msi.TranscodingURL,
		}
	}
	return pbms
}

func item2pb(item *emby.Items) *pb.Item {
	pi := &pb.Item{
		Id:              item.ID,
		Name:            item.Name,
		Type:            item.Type,
		IsFolder:        item.IsFolder,
		ParentId:        item.ParentID,
		SeasonName:      item.SeasonName,
		SeasonId:        item.SeasonID,
		SeriesName:      item.SeriesName,
		SeriesId:        item.SeriesID,
		MediaSourceInfo: mediaSources2pb(item.MediaSources),
		CollectionType:  item.CollectionType,
	}
	return pi
}

func (a *EmbyService) GetItems(ctx context.Context, req *pb.GetItemsReq) (*pb.GetItemsResp, error) {
	cli := emby.NewClient(req.Host, emby.WithContext(ctx), emby.WithKey(req.Token))
	opts := []emby.GetItemsOptionFunc{
		emby.WithSortBy("SortName"),
		emby.WithSortOrderAsc(),
	}
	if req.SearchTerm != "" {
		opts = append(opts,
			emby.WithSearch(req.SearchTerm),
			emby.WithRecursive(),
		)
	} else {
		opts = append(opts,
			emby.WithParentId(req.ParentId),
		)
	}
	r, err := cli.GetItems(opts...)
	if err != nil {
		return nil, err
	}
	var items []*pb.Item
	for _, item := range r.Items {
		items = append(items, item2pb(&item))
	}
	return &pb.GetItemsResp{
		Items:            items,
		TotalRecordCount: r.TotalRecordCount,
	}, nil
}

func (a *EmbyService) GetItem(ctx context.Context, req *pb.GetItemReq) (*pb.Item, error) {
	cli := emby.NewClient(req.Host, emby.WithContext(ctx), emby.WithKey(req.Token))
	r, err := cli.GetItem(req.ItemId)
	if err != nil {
		return nil, err
	}
	return item2pb(r), nil
}

func (a *EmbyService) FsList(ctx context.Context, req *pb.FsListReq) (*pb.FsListResp, error) {
	cli := emby.NewClient(req.Host, emby.WithContext(ctx), emby.WithKey(req.Token), emby.WithUserID(req.UserId))
	var resp *emby.ItemsResp
	var paths []*pb.Path
	var err error
	if req.SearchTerm != "" {
		opts := []emby.GetItemsOptionFunc{
			emby.WithSortBy("SortName"),
			emby.WithSortOrderAsc(),
			emby.WithRecursive(),
			emby.WithSearch(req.SearchTerm),
		}
		if req.StartIndex != 0 || req.Limit != 0 {
			opts = append(opts,
				emby.WithStartIndex(req.StartIndex),
				emby.WithLimit(req.Limit),
			)
		}
		if req.Path != "" {
			opts = append(opts, emby.WithParentId(req.Path))
			paths, err = genPath(cli, req.Path, nil)
			if err != nil {
				return nil, err
			}
		} else {
			paths = []*pb.Path{
				{
					Name: "Home",
					Path: "",
				},
			}
		}
		resp, err = cli.UserItems(opts...)
	} else if req.Path == "" {
		resp, err = cli.UserViews()
		paths = []*pb.Path{
			{
				Name: "Home",
				Path: "",
			},
		}
	} else {
		var item *emby.Items
		item, err = cli.UserItemsByID(req.Path)
		if err != nil {
			return nil, err
		}
		switch item.Type {
		case "CollectionFolder":
			opts := []emby.GetItemsOptionFunc{
				emby.WithSortBy("SortName"),
				emby.WithSortOrderAsc(),
				emby.WithParentId(item.ID),
				emby.WithRecursive(),
			}
			if req.StartIndex != 0 || req.Limit != 0 {
				opts = append(opts,
					emby.WithStartIndex(req.StartIndex),
					emby.WithLimit(req.Limit),
				)
			}
			switch item.CollectionType {
			case "movies":
				opts = append(opts,
					emby.WithIncludeItemTypes("Movie"),
				)
			case "tvshows":
				opts = append(opts,
					emby.WithIncludeItemTypes("Series"),
				)
			}
			resp, err = cli.UserItems(opts...)
		case "Series":
			resp, err = cli.Seasons(item.ID)
		case "Season":
			resp, err = cli.Episodes(item.SeriesID, item.ID)
		default:
			return nil, fmt.Errorf("unknown type: %s", item.Type)
		}
		if err != nil {
			return nil, err
		}
		paths, err = genPath(cli, item.ID, map[string]*emby.Items{
			item.ID: item,
		})
	}
	if err != nil {
		return nil, err
	}
	var items []*pb.Item
	for _, item := range resp.Items {
		items = append(items, item2pb(&item))
	}
	return &pb.FsListResp{
		Items: items,
		Paths: paths,
		Total: resp.TotalRecordCount,
	}, nil
}

const genPathMaxDepth = 5

func genPath(cli *emby.Client, id string, caches map[string]*emby.Items) ([]*pb.Path, error) {
	var paths []*pb.Path
	var (
		item  *emby.Items
		err   error
		depth int
	)
	for {
		if depth > genPathMaxDepth {
			return nil, errors.New("genPath: max depth reached")
		}
		if len(caches) > 0 {
			var ok bool
			if item, ok = caches[id]; !ok {
				item, err = cli.UserItemsByID(id)
				if err != nil {
					return nil, err
				}
			}
		} else {
			item, err = cli.UserItemsByID(id)
			if err != nil {
				return nil, err
			}
		}
		depth++
		switch item.Type {
		case "Series", "Folder":
			i, err := strconv.Atoi(item.ParentID)
			if err != nil {
				return nil, err
			}
			id = strconv.Itoa(i + 1)
		case "UserRootFolder":
			return append([]*pb.Path{
				{
					Name: "Home",
					Path: "",
				},
			}, paths...), nil
		default:
			id = item.ParentID
		}
		paths = append([]*pb.Path{
			{
				Name: item.Name,
				Path: item.ID,
			},
		}, paths...)
	}
}

func (a *EmbyService) GetSystemInfo(ctx context.Context, req *pb.SystemInfoReq) (*pb.SystemInfoResp, error) {
	cli := emby.NewClient(req.Host, emby.WithContext(ctx), emby.WithKey(req.Token))
	r, err := cli.SystemInfo()
	if err != nil {
		return nil, err
	}
	return &pb.SystemInfoResp{
		SystemUpdateLevel:                    r.SystemUpdateLevel,
		OperatingSystemDisplayName:           r.OperatingSystemDisplayName,
		PackageName:                          r.PackageName,
		HasPendingRestart:                    r.HasPendingRestart,
		IsShuttingDown:                       r.IsShuttingDown,
		SupportsLibraryMonitor:               r.SupportsLibraryMonitor,
		WebSocketPortNumber:                  r.WebSocketPortNumber,
		CanSelfRestart:                       r.CanSelfRestart,
		CanSelfUpdate:                        r.CanSelfUpdate,
		CanLaunchWebBrowser:                  r.CanLaunchWebBrowser,
		ProgramDataPath:                      r.ProgramDataPath,
		ItemsByNamePath:                      r.ItemsByNamePath,
		CachePath:                            r.CachePath,
		LogPath:                              r.LogPath,
		InternalMetadataPath:                 r.InternalMetadataPath,
		TranscodingTempPath:                  r.TranscodingTempPath,
		HttpServerPortNumber:                 r.HttpServerPortNumber,
		SupportsHttps:                        r.SupportsHttps,
		HttpsPortNumber:                      r.HttpsPortNumber,
		HasUpdateAvailable:                   r.HasUpdateAvailable,
		SupportsAutoRunAtStartup:             r.SupportsAutoRunAtStartup,
		HardwareAccelerationRequiresPremiere: r.HardwareAccelerationRequiresPremiere,
		LocalAddress:                         r.LocalAddress,
		WanAddress:                           r.WanAddress,
		ServerName:                           r.ServerName,
		Version:                              r.Version,
		OperatingSystem:                      r.OperatingSystem,
		Id:                                   r.Id,
	}, nil
}

func (a *EmbyService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.Empty, error) {
	err := emby.NewClient(req.Host, emby.WithContext(ctx), emby.WithKey(req.Token)).Logout()
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (a *EmbyService) PlaybackInfo(ctx context.Context, req *pb.PlaybackInfoReq) (*pb.PlaybackInfoResp, error) {
	cli := emby.NewClient(req.Host, emby.WithContext(ctx), emby.WithKey(req.Token), emby.WithUserID(req.UserId))
	r, err := cli.UserPlaybackInfo(req.ItemId)
	if err != nil {
		return nil, err
	}
	return &pb.PlaybackInfoResp{
		PlaySessionID:   r.PlaySessionID,
		MediaSourceInfo: mediaSources2pb(r.MediaSources),
	}, nil
}
