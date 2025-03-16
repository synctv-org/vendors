package alist

import (
	"context"

	pb "github.com/synctv-org/vendors/api/alist"
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/vendors/alist"
)

type AlistService struct {
	pb.UnimplementedAlistServer
}

func NewAlistService(c *conf.AlistConfig) *AlistService {
	return &AlistService{}
}

func (a *AlistService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	opts := []alist.LoginOpt{}
	if req.GetHashed() {
		opts = append(opts, alist.WithHashed())
	}
	s, err := alist.Login(ctx, req.GetHost(), alist.LoginReq{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		Token: s,
	}, nil
}

func (a *AlistService) Me(ctx context.Context, req *pb.MeReq) (*pb.MeResp, error) {
	cli, err := alist.NewClient(req.GetHost(), req.GetToken(), alist.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	r, err := cli.Me()
	if err != nil {
		return nil, err
	}
	return &pb.MeResp{
		Id:         r.ID,
		Username:   r.Username,
		Password:   r.Password,
		BasePath:   r.BasePath,
		Role:       r.Role,
		Disabled:   r.Disabled,
		Permission: r.Permission,
		SsoId:      r.SsoID,
		Otp:        r.Otp,
	}, nil
}

func related2pb(r *alist.FsGetRelated) *pb.FsGetResp_FsGetRelated {
	return &pb.FsGetResp_FsGetRelated{
		Name:     r.Name,
		Created:  uint64(r.Created.UnixMilli()),
		Modified: uint64(r.Modified.UnixMilli()),
		Sign:     r.Sign,
		Thumb:    r.Thumb,
		Hashinfo: r.Hashinfo,
		Size:     r.Size,
		IsDir:    r.IsDir,
		Type:     r.Type,
	}
}

func relateds2pb(rs []*alist.FsGetRelated) []*pb.FsGetResp_FsGetRelated {
	res := make([]*pb.FsGetResp_FsGetRelated, len(rs))
	for i, r := range rs {
		res[i] = related2pb(r)
	}
	return res
}

func (a *AlistService) FsGet(ctx context.Context, req *pb.FsGetReq) (*pb.FsGetResp, error) {
	cli, err := alist.NewClient(req.GetHost(), req.GetToken(), alist.WithContext(ctx), alist.WithUserAgent(req.GetUserAgent()))
	if err != nil {
		return nil, err
	}
	r, err := cli.FsGet(&alist.FsGetReq{
		Path:     req.GetPath(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.FsGetResp{
		Name:     r.Name,
		Size:     r.Size,
		IsDir:    r.IsDir,
		Modified: uint64(r.Modified.UnixMilli()),
		Created:  uint64(r.Created.UnixMilli()),
		Sign:     r.Sign,
		Thumb:    r.Thumb,
		Type:     r.Type,
		Hashinfo: r.Hashinfo,
		RawUrl:   r.RawURL,
		Readme:   r.Readme,
		Provider: r.Provider,
		Related:  relateds2pb(r.Related),
	}, nil
}

func (a *AlistService) FsList(ctx context.Context, req *pb.FsListReq) (*pb.FsListResp, error) {
	cli, err := alist.NewClient(req.GetHost(), req.GetToken(), alist.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	r, err := cli.FsList(&alist.FsListReq{
		Path:     req.GetPath(),
		Password: req.GetPassword(),
		Page:     req.GetPage(),
		PerPage:  req.GetPerPage(),
		Refresh:  req.GetRefresh(),
	})
	if err != nil {
		return nil, err
	}
	content := make([]*pb.FsListResp_FsListContent, len(r.Content))
	for i, v := range r.Content {
		content[i] = &pb.FsListResp_FsListContent{
			Name:     v.Name,
			Size:     v.Size,
			IsDir:    v.IsDir,
			Modified: uint64(v.Modified.UnixMilli()),
			Sign:     v.Sign,
			Thumb:    v.Thumb,
			Type:     v.Type,
		}
	}
	return &pb.FsListResp{
		Content:  content,
		Total:    r.Total,
		Readme:   r.Readme,
		Write:    r.Write,
		Provider: r.Provider,
	}, nil
}

func (a *AlistService) FsOther(ctx context.Context, req *pb.FsOtherReq) (*pb.FsOtherResp, error) {
	cli, err := alist.NewClient(req.GetHost(), req.GetToken(), alist.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	r, err := cli.FsOther(&alist.FsOtherReq{
		Path:     req.GetPath(),
		Method:   req.GetMethod(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	videoPreviewPlayInfo := pb.FsOtherResp_VideoPreviewPlayInfo{
		Category:                        r.VideoPreviewPlayInfo.Category,
		LiveTranscodingSubtitleTaskList: make([]*pb.FsOtherResp_VideoPreviewPlayInfo_LiveTranscodingSubtitleTaskList, len(r.VideoPreviewPlayInfo.LiveTranscodingSubtitleTaskList)),
		LiveTranscodingTaskList:         make([]*pb.FsOtherResp_VideoPreviewPlayInfo_LiveTranscodingTaskList, len(r.VideoPreviewPlayInfo.LiveTranscodingTaskList)),
		Meta: &pb.FsOtherResp_VideoPreviewPlayInfo_Meta{
			Duration: r.VideoPreviewPlayInfo.Meta.Duration,
			Height:   r.VideoPreviewPlayInfo.Meta.Height,
			Width:    r.VideoPreviewPlayInfo.Meta.Width,
		},
	}

	for i, v := range r.VideoPreviewPlayInfo.LiveTranscodingSubtitleTaskList {
		videoPreviewPlayInfo.LiveTranscodingSubtitleTaskList[i] = &pb.FsOtherResp_VideoPreviewPlayInfo_LiveTranscodingSubtitleTaskList{
			Language: v.Language,
			Status:   v.Status,
			Url:      v.URL,
		}
	}

	for i, v := range r.VideoPreviewPlayInfo.LiveTranscodingTaskList {
		videoPreviewPlayInfo.LiveTranscodingTaskList[i] = &pb.FsOtherResp_VideoPreviewPlayInfo_LiveTranscodingTaskList{
			Stage:          v.Stage,
			Status:         v.Status,
			TemplateHeight: v.TemplateHeight,
			TemplateId:     v.TemplateID,
			TemplateName:   v.TemplateName,
			TemplateWidth:  v.TemplateWidth,
			Url:            v.URL,
		}
	}

	return &pb.FsOtherResp{
		DriveId:              r.DriveID,
		FileId:               r.FileID,
		VideoPreviewPlayInfo: &videoPreviewPlayInfo,
	}, nil
}

func (a *AlistService) FsSearch(ctx context.Context, req *pb.FsSearchReq) (*pb.FsSearchResp, error) {
	cli, err := alist.NewClient(req.GetHost(), req.GetToken(), alist.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	r, err := cli.FsSearch(&alist.FsSearchReq{
		Parent:   req.GetParent(),
		Password: req.GetPassword(),
		Keywords: req.GetKeywords(),
		Scope:    req.GetScope(),
		Page:     req.GetPage(),
		PerPage:  req.GetPerPage(),
	})
	if err != nil {
		return nil, err
	}
	content := make([]*pb.FsSearchResp_FsSearchContent, len(r.Content))
	for i, v := range r.Content {
		content[i] = &pb.FsSearchResp_FsSearchContent{
			Parent: v.Parent,
			Name:   v.Name,
			Size:   v.Size,
			IsDir:  v.IsDir,
			Type:   v.Type,
		}
	}
	return &pb.FsSearchResp{
		Content: content,
		Total:   r.Total,
	}, nil
}
