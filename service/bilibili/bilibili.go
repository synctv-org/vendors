package bilibili

import (
	"context"

	pb "github.com/synctv-org/vendors/api/bilibili"
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/utils"
	"github.com/synctv-org/vendors/vendors/bilibili"
)

type BilibiliService struct {
	pb.UnimplementedBilibiliServer
}

func NewBilibiliService(c *conf.BilibiliConfig) *BilibiliService {
	return &BilibiliService{}
}

func (s *BilibiliService) NewQRCode(ctx context.Context, req *pb.Empty) (*pb.NewQRCodeResp, error) {
	r, err := bilibili.NewQRCode(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.NewQRCodeResp{
		Url: r.URL,
		Key: r.Key,
	}, nil
}

func (s *BilibiliService) LoginWithQRCode(ctx context.Context, req *pb.LoginWithQRCodeReq) (*pb.LoginWithQRCodeResp, error) {
	status, c, err := bilibili.LoginWithQRCode(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	resp := &pb.LoginWithQRCodeResp{
		Status: status,
	}
	if c != nil {
		resp.Cookies = map[string]string{c.Name: c.Value}
	}
	return resp, nil
}

func (s *BilibiliService) NewCaptcha(ctx context.Context, req *pb.Empty) (*pb.NewCaptchaResp, error) {
	cr, err := bilibili.NewCaptcha(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.NewCaptchaResp{
		Token:     cr.Token,
		Gt:        cr.Gt,
		Challenge: cr.Challenge,
	}, nil
}

func (s *BilibiliService) NewSMS(ctx context.Context, req *pb.NewSMSReq) (*pb.NewSMSResp, error) {
	captchaKey, err := bilibili.NewSMS(ctx, req.Phone, req.Token, req.Challenge, req.Validate)
	if err != nil {
		return nil, err
	}
	return &pb.NewSMSResp{
		CaptchaKey: captchaKey,
	}, nil
}

func (s *BilibiliService) LoginWithSMS(ctx context.Context, req *pb.LoginWithSMSReq) (*pb.LoginWithSMSResp, error) {
	c, err := bilibili.LoginWithSMS(ctx, req.Phone, req.Code, req.CaptchaKey)
	if err != nil {
		return nil, err
	}
	return &pb.LoginWithSMSResp{
		Cookies: map[string]string{c.Name: c.Value},
	}, nil
}

func (s *BilibiliService) ParseVideoPage(ctx context.Context, req *pb.ParseVideoPageReq) (*pb.VideoPageInfo, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.Cookies), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	r, err := c.ParseVideoPage(req.Aid, req.Bvid)
	if err != nil {
		return nil, err
	}
	resp := &pb.VideoPageInfo{
		Title:  r.Title,
		Actors: r.Actors,
	}
	if len(r.VideoInfos) != 0 {
		resp.VideoInfos = make([]*pb.VideoInfo, len(r.VideoInfos))
		for i, v := range r.VideoInfos {
			resp.VideoInfos[i] = &pb.VideoInfo{
				Bvid:       v.Bvid,
				Cid:        v.Cid,
				Epid:       v.Epid,
				Name:       v.Name,
				CoverImage: v.CoverImage,
			}
		}
	}
	return resp, nil
}

func (s *BilibiliService) GetVideoURL(ctx context.Context, req *pb.GetVideoURLReq) (*pb.VideoURL, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.Cookies), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	conf := []bilibili.GetVideoURLConfig{}
	if req.Quality != 0 {
		conf = append(conf, bilibili.WithQuality(req.Quality))
	}
	r, err := c.GetVideoURL(req.Aid, req.Bvid, req.Cid, conf...)
	if err != nil {
		return nil, err
	}
	return &pb.VideoURL{
		AcceptDescription: r.AcceptDescription,
		AcceptQuality:     r.AcceptQuality,
		CurrentQuality:    r.CurrentQuality,
		Url:               r.URL,
	}, nil
}

func (s *BilibiliService) GetDashVideoURL(ctx context.Context, req *pb.GetDashVideoURLReq) (*pb.GetDashVideoURLResp, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.Cookies), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	opt := make([]bilibili.GetDashVideoURLConfig, 0)
	m, m2, err := c.GetDashVideoURL(req.Aid, req.Bvid, req.Cid, opt...)
	if err != nil {
		return nil, err
	}
	str, err := m.WriteToString()
	if err != nil {
		return nil, err
	}
	str2, err := m2.WriteToString()
	if err != nil {
		return nil, err
	}
	return &pb.GetDashVideoURLResp{
		Mpd:     str,
		HevcMpd: str2,
	}, nil
}

func (s *BilibiliService) GetSubtitles(ctx context.Context, req *pb.GetSubtitlesReq) (*pb.GetSubtitlesResp, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.Cookies), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	sub, err := c.GetSubtitles(req.Aid, req.Bvid, req.Cid)
	if err != nil {
		return nil, err
	}
	resp := &pb.GetSubtitlesResp{}
	if len(sub) != 0 {
		resp.Subtitles = make(map[string]string, len(sub))
		for _, v := range sub {
			resp.Subtitles[v.Name] = v.URL
		}
	}
	return resp, nil
}

func (s *BilibiliService) ParsePGCPage(ctx context.Context, req *pb.ParsePGCPageReq) (*pb.VideoPageInfo, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.Cookies), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	r, err := c.ParsePGCPage(req.Epid, req.Ssid)
	if err != nil {
		return nil, err
	}
	resp := &pb.VideoPageInfo{
		Title:  r.Title,
		Actors: r.Actors,
	}
	if len(r.VideoInfos) != 0 {
		resp.VideoInfos = make([]*pb.VideoInfo, len(r.VideoInfos))
		for i, v := range r.VideoInfos {
			resp.VideoInfos[i] = &pb.VideoInfo{
				Bvid:       v.Bvid,
				Cid:        v.Cid,
				Epid:       v.Epid,
				Name:       v.Name,
				CoverImage: v.CoverImage,
			}
		}
	}
	return resp, nil
}

func (s *BilibiliService) GetPGCURL(ctx context.Context, req *pb.GetPGCURLReq) (*pb.VideoURL, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.Cookies), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	conf := []bilibili.GetVideoURLConfig{}
	if req.Quality != 0 {
		conf = append(conf, bilibili.WithQuality(req.Quality))
	}
	r, err := c.GetPGCURL(req.Epid, req.Cid, conf...)
	if err != nil {
		return nil, err
	}
	return &pb.VideoURL{
		AcceptDescription: r.AcceptDescription,
		AcceptQuality:     r.AcceptQuality,
		CurrentQuality:    r.CurrentQuality,
		Url:               r.URL,
	}, nil
}

func (s *BilibiliService) GetDashPGCURL(ctx context.Context, req *pb.GetDashPGCURLReq) (*pb.GetDashPGCURLResp, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.Cookies), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	opt := make([]bilibili.GetDashVideoURLConfig, 0)
	m, m2, err := c.GetDashPGCURL(req.Epid, req.Cid, opt...)
	if err != nil {
		return nil, err
	}
	str, err := m.WriteToString()
	if err != nil {
		return nil, err
	}
	str2, err := m2.WriteToString()
	if err != nil {
		return nil, err
	}
	return &pb.GetDashPGCURLResp{
		Mpd:     str,
		HevcMpd: str2,
	}, nil
}

func (s *BilibiliService) UserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.Cookies), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	info, err := c.UserInfo()
	if err != nil {
		return nil, err
	}
	return &pb.UserInfoResp{
		IsLogin:  info.IsLogin,
		Username: info.Username,
		Face:     info.Face,
		IsVip:    info.IsVip,
	}, nil
}

func (s *BilibiliService) Match(ctx context.Context, req *pb.MatchReq) (*pb.MatchResp, error) {
	t, id, err := bilibili.Match(req.Url)
	if err != nil {
		return nil, err
	}
	return &pb.MatchResp{
		Type: t,
		Id:   id,
	}, nil
}
