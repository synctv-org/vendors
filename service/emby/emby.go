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

func (a *EmbyService) GetItems(ctx context.Context, req *pb.GetItemsReq) (*pb.GetItemsResp, error) {
	cli := emby.NewClient(req.Host, emby.WithContext(ctx), emby.WithKey(req.Token))
	r, err := cli.GetItems(req.ParentId)
	if err != nil {
		return nil, err
	}
	var items []*pb.Item
	for _, item := range r.Items {
		items = append(items, &pb.Item{
			Id:         item.Id,
			Name:       item.Name,
			Type:       item.Type,
			IsFolder:   item.IsFolder,
			ParentId:   item.ParentId,
			SeasonName: item.SeasonName,
			SeasonId:   item.SeasonId,
			SeriesName: item.SeriesName,
			SeriesId:   item.SeriesId,
		})
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
	return &pb.Item{
		Id:         r.Id,
		Name:       r.Name,
		Type:       r.Type,
		IsFolder:   r.IsFolder,
		ParentId:   r.ParentId,
		SeasonName: r.SeasonName,
		SeasonId:   r.SeasonId,
		SeriesName: r.SeriesName,
		SeriesId:   r.SeriesId,
	}, nil
}

func (a *EmbyService) FsList(ctx context.Context, req *pb.FsListReq) (*pb.FsListResp, error) {
	cli := emby.NewClient(req.Host, emby.WithContext(ctx), emby.WithKey(req.Token))
	opts := make([]emby.GetItemsOptionFunc, 0)
	if req.StartIndex != 0 {
		opts = append(opts, emby.WithStartIndex(req.StartIndex))
	}
	if req.Limit != 0 {
		opts = append(opts, emby.WithLimit(req.Limit))
	}
	r, err := cli.GetItems(req.Path, opts...)
	if err != nil {
		return nil, err
	}
	var items []*pb.Item
	for _, item := range r.Items {
		items = append(items, &pb.Item{
			Id:         item.Id,
			Name:       item.Name,
			Type:       item.Type,
			IsFolder:   item.IsFolder,
			ParentId:   item.ParentId,
			SeasonName: item.SeasonName,
			SeasonId:   item.SeasonId,
			SeriesName: item.SeriesName,
			SeriesId:   item.SeriesId,
		})
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
				Path: item.Id,
			},
		}, paths...)
		switch item.Type {
		case "Series", "Folder":
			i, err := strconv.Atoi(item.ParentId)
			if err != nil {
				return nil, err
			}
			id = strconv.Itoa(i + 1)
		default:
			id = item.ParentId
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
