package emby

import (
	"context"

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
