package emby

import (
	"context"
	"fmt"

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
	cli := emby.NewClient(req.GetHost(), emby.WithContext(ctx))
	r, err := cli.GetAPIKey(req.GetUsername(), req.GetPassword())
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
	cli := emby.NewClient(req.GetHost(), emby.WithContext(ctx), emby.WithKey(req.GetToken()), emby.WithUserID(req.GetUserId()))
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
	pbmsi := make([]*pb.MediaStreamInfo, len(msi))
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
	pbms := make([]*pb.MediaSourceInfo, len(ms))
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
	cli := emby.NewClient(req.GetHost(), emby.WithContext(ctx), emby.WithKey(req.GetToken()), emby.WithUserID(req.GetUserId()))
	opts := []emby.QueryFunc{
		emby.WithSortBy("SortName"),
		emby.WithSortOrderAsc(),
		emby.WithNotFolder(),
	}
	if req.GetSearchTerm() != "" {
		opts = append(opts,
			emby.WithSearch(req.GetSearchTerm()),
			emby.WithRecursive(),
		)
	} else {
		opts = append(opts,
			emby.WithParentID(req.GetParentId()),
		)
	}
	r, err := cli.UserItems(opts...)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.Item, len(r.Items))
	for i, item := range r.Items {
		items[i] = item2pb(item)
	}
	return &pb.GetItemsResp{
		Items:            items,
		TotalRecordCount: r.TotalRecordCount,
	}, nil
}

func (a *EmbyService) GetItem(ctx context.Context, req *pb.GetItemReq) (*pb.Item, error) {
	cli := emby.NewClient(req.GetHost(), emby.WithContext(ctx), emby.WithKey(req.GetToken()))
	r, err := cli.GetItem(req.GetItemId())
	if err != nil {
		return nil, err
	}
	return item2pb(r), nil
}

func (a *EmbyService) FsList(ctx context.Context, req *pb.FsListReq) (*pb.FsListResp, error) {
	cli := emby.NewClient(req.GetHost(), emby.WithContext(ctx), emby.WithKey(req.GetToken()), emby.WithUserID(req.GetUserId()))
	var resp *emby.ItemsResp
	var paths []*pb.Path
	var err error
	opts := []emby.QueryFunc{}
	if req.GetStartIndex() != 0 || req.GetLimit() != 0 {
		opts = append(opts,
			emby.WithStartIndex(req.GetStartIndex()),
			emby.WithLimit(req.GetLimit()),
		)
	}
	if req.GetSearchTerm() != "" {
		opts = append(opts,
			emby.WithSortBy("SortName"),
			emby.WithSortOrderAsc(),
			emby.WithRecursive(),
			emby.WithSearch(req.GetSearchTerm()),
		)
		if req.GetPath() != "" {
			opts = append(opts, emby.WithParentID(req.GetPath()))
			var item *emby.Items
			item, err = cli.UserItemsByID(req.GetPath())
			if err != nil {
				return nil, err
			}
			paths = []*pb.Path{
				{
					Name: "Home",
					Path: "",
				},
				{
					Name: item.Name,
					Path: item.ID,
				},
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
	} else if req.GetPath() == "" {
		resp, err = cli.UserViews(opts...)
		paths = []*pb.Path{
			{
				Name: "Home",
				Path: "",
			},
		}
	} else {
		var item *emby.Items
		item, err = cli.UserItemsByID(req.GetPath())
		if err != nil {
			return nil, err
		}
		switch item.Type {
		case "CollectionFolder":
			opts = append(opts,
				emby.WithSortBy("SortName"),
				emby.WithSortOrderAsc(),
				emby.WithParentID(item.ID),
				emby.WithRecursive(),
			)
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
			resp, err = cli.Seasons(item.ID, opts...)
		case "Season":
			resp, err = cli.Seasons(item.ParentID,
				emby.WithStartIndex(req.GetStartIndex()),
				emby.WithLimit(req.GetLimit()),
			)
			if err != nil {
				return nil, err
			}
			if resp.TotalRecordCount == 1 {
				resp, err = cli.UserItems(
					emby.WithParentID(item.ID),
					emby.WithRecursive(),
				)
			} else {
				resp, err = cli.Episodes(item.SeriesID, item.ID, opts...)
			}
		default:
			return nil, fmt.Errorf("unknown type: %s", item.Type)
		}
		paths = []*pb.Path{
			{
				Name: "Home",
				Path: "",
			},
			{
				Name: item.Name,
				Path: item.ID,
			},
		}
	}
	if err != nil {
		return nil, err
	}
	items := make([]*pb.Item, len(resp.Items))
	for i, item := range resp.Items {
		items[i] = item2pb(item)
	}
	return &pb.FsListResp{
		Items: items,
		Paths: paths,
		Total: resp.TotalRecordCount,
	}, nil
}

func (a *EmbyService) GetSystemInfo(ctx context.Context, req *pb.SystemInfoReq) (*pb.SystemInfoResp, error) {
	cli := emby.NewClient(req.GetHost(), emby.WithContext(ctx), emby.WithKey(req.GetToken()))
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
		HttpServerPortNumber:                 r.HTTPServerPortNumber,
		SupportsHttps:                        r.SupportsHTTPS,
		HttpsPortNumber:                      r.HTTPSPortNumber,
		HasUpdateAvailable:                   r.HasUpdateAvailable,
		SupportsAutoRunAtStartup:             r.SupportsAutoRunAtStartup,
		HardwareAccelerationRequiresPremiere: r.HardwareAccelerationRequiresPremiere,
		LocalAddress:                         r.LocalAddress,
		WanAddress:                           r.WanAddress,
		ServerName:                           r.ServerName,
		Version:                              r.Version,
		OperatingSystem:                      r.OperatingSystem,
		Id:                                   r.ID,
	}, nil
}

func (a *EmbyService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.Empty, error) {
	err := emby.NewClient(req.GetHost(), emby.WithContext(ctx), emby.WithKey(req.GetToken())).Logout()
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (a *EmbyService) PlaybackInfo(ctx context.Context, req *pb.PlaybackInfoReq) (*pb.PlaybackInfoResp, error) {
	cli := emby.NewClient(req.GetHost(), emby.WithContext(ctx), emby.WithKey(req.GetToken()), emby.WithUserID(req.GetUserId()))
	opts := []emby.UserPlaybackInfoOption{}
	if req.GetMediaSourceId() != "" {
		opts = append(opts, emby.WithMediaSourceID(req.GetMediaSourceId()))
	}
	if req.GetSubtitleStreamIndex() != 0 {
		opts = append(opts, emby.WithSubtitleStreamIndex(int(req.GetSubtitleStreamIndex())))
	}
	if req.GetAudioStreamIndex() != 0 {
		opts = append(opts, emby.WithAudioStreamIndex(int(req.GetAudioStreamIndex())))
	}
	if req.GetMaxStreamingBitrate() != 0 {
		opts = append(opts, emby.WithMaxStreamingBitrate(int(req.GetMaxStreamingBitrate())))
	}
	r, err := cli.UserPlaybackInfo(req.GetItemId(), opts...)
	if err != nil {
		return nil, err
	}
	return &pb.PlaybackInfoResp{
		PlaySessionID:   r.PlaySessionID,
		MediaSourceInfo: mediaSources2pb(r.MediaSources),
	}, nil
}

func (a *EmbyService) DeleteActiveEncodeings(ctx context.Context, req *pb.DeleteActiveEncodeingsReq) (*pb.Empty, error) {
	cli := emby.NewClient(req.GetHost(), emby.WithContext(ctx), emby.WithKey(req.GetToken()))
	return nil, cli.DeleteActiveEncodeings(req.GetPalySessionId())
}
