package webdav

import (
	"context"
	"errors"

	"github.com/studio-b12/gowebdav"
	pb "github.com/synctv-org/vendors/api/webdav"
	"github.com/synctv-org/vendors/conf"
	"github.com/zijiren233/go-uhc"
)

type WebdavService struct {
	pb.UnimplementedWebdavServer
}

func NewWebdavService(c *conf.WebdavConfig) *WebdavService {
	return &WebdavService{}
}

func (a *WebdavService) FsTest(ctx context.Context, req *pb.FsTestReq) (*pb.Empty, error) {
	cli := gowebdav.NewClient(req.Host, req.Username, req.Password)
	cli.SetTransport(uhc.DefaultTransport)
	return nil, cli.Connect()
}

func (a *WebdavService) FsGet(ctx context.Context, req *pb.FsGetReq) (*pb.FsGetResp, error) {
	cli := gowebdav.NewClient(req.Host, req.Username, req.Password)
	cli.SetTransport(uhc.DefaultTransport)
	info, err := cli.Stat(req.Path)
	if err != nil {
		return nil, err
	}
	return &pb.FsGetResp{
		Name:     info.Name(),
		Size:     info.Size(),
		IsDir:    info.IsDir(),
		Modified: info.ModTime().UnixMilli(),
	}, nil
}

func (a *WebdavService) FsList(ctx context.Context, req *pb.FsListReq) (*pb.FsListResp, error) {
	cli := gowebdav.NewClient(req.Host, req.Username, req.Password)
	cli.SetTransport(uhc.DefaultTransport)
	infos, err := cli.ReadDir(req.Path)
	if err != nil {
		return nil, err
	}
	var items []*pb.FsListResp_FsListContent = make([]*pb.FsListResp_FsListContent, 0, req.PerPage)
	var low, high int
	low, high = int(req.Page-1)*int(req.PerPage), int(req.Page)*int(req.PerPage)
	if low > len(infos) {
		low = len(infos)
	}
	if high > len(infos) {
		high = len(infos)
	}
	for _, info := range infos[low:high] {
		items = append(items, &pb.FsListResp_FsListContent{
			Name:     info.Name(),
			Size:     info.Size(),
			IsDir:    info.IsDir(),
			Modified: info.ModTime().UnixMilli(),
		})
	}
	return &pb.FsListResp{
		Content: items,
		Total:   uint64(len(infos)),
	}, nil
}

func (a *WebdavService) FsSearch(ctx context.Context, req *pb.FsSearchReq) (*pb.FsSearchResp, error) {
	return nil, errors.New("not implemented")
}
