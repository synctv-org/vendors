package emby

import (
	"context"
	"errors"
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
		Token:  r.AccessToken,
		UserId: r.UserID,
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
		MediaSourceInfo: make([]*pb.MediaSourceInfo, len(item.MediaSources)),
	}
	for i, msi := range item.MediaSources {
		pi.MediaSourceInfo[i] = &pb.MediaSourceInfo{
			Id:                         msi.ID,
			Name:                       msi.Name,
			Path:                       msi.Path,
			Protocol:                   msi.Protocol,
			Container:                  msi.Container,
			DefaultSubtitleStreamIndex: msi.DefaultSubtitleStreamIndex,
			DefaultAudioStreamIndex:    msi.DefaultAudioStreamIndex,
			MediaStreamInfo:            make([]*pb.MediaStreamInfo, len(msi.MediaStreams)),
		}
		for j, msi := range msi.MediaStreams {
			pi.MediaSourceInfo[i].MediaStreamInfo[j] = &pb.MediaStreamInfo{
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
	}
	return pi
}

func (a *EmbyService) GetItems(ctx context.Context, req *pb.GetItemsReq) (*pb.GetItemsResp, error) {
	cli := emby.NewClient(req.Host, emby.WithContext(ctx), emby.WithKey(req.Token))
	opts := []emby.GetItemsOptionFunc{
		emby.WithParentId(req.ParentId),
	}
	if req.SearchTerm != "" {
		opts = append(opts, emby.WithRecursive(), emby.WithSortBy("SortName"), emby.WithSortOrderAsc(), emby.WithSearch(req.SearchTerm))
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
	cli := emby.NewClient(req.Host, emby.WithContext(ctx), emby.WithKey(req.Token))
	opts := []emby.GetItemsOptionFunc{
		emby.WithParentId(req.Path),
	}
	if req.SearchTerm != "" {
		opts = append(opts, emby.WithRecursive(), emby.WithSortBy("SortName"), emby.WithSortOrderAsc(), emby.WithSearch(req.SearchTerm))
	}
	if req.StartIndex != 0 {
		opts = append(opts, emby.WithStartIndex(req.StartIndex))
	}
	if req.Limit != 0 {
		opts = append(opts, emby.WithLimit(req.Limit))
	}
	r, err := cli.GetItems(opts...)
	if err != nil {
		return nil, err
	}
	var items []*pb.Item
	for _, item := range r.Items {
		items = append(items, item2pb(&item))
	}
	paths, err := genPath(cli, req.Path)
	if err != nil {
		return nil, err
	}
	return &pb.FsListResp{
		Items: items,
		Paths: paths,
		Total: r.TotalRecordCount,
	}, nil
}

func genPath(cli *emby.Client, id string) ([]*pb.Path, error) {
	var paths []*pb.Path
	for {
		if id == "1" || id == "" {
			break
		}
		item, err := cli.GetItem(id)
		if err != nil {
			return nil, err
		}
		if !item.IsFolder {
			return nil, errors.New("not a folder")
		}
		paths = append([]*pb.Path{
			{
				Name: item.Name,
				Path: item.ID,
			},
		}, paths...)
		switch item.Type {
		case "Series", "Folder":
			i, err := strconv.Atoi(item.ParentID)
			if err != nil {
				return nil, err
			}
			id = strconv.Itoa(i + 1)
		default:
			id = item.ParentID
		}
	}
	return append([]*pb.Path{
		{
			Name: "Home",
			Path: "1",
		},
	}, paths...), nil
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
