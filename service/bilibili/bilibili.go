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
	status, c, err := bilibili.LoginWithQRCode(ctx, req.GetKey())
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
	captchaKey, err := bilibili.NewSMS(ctx, req.GetPhone(), req.GetToken(), req.GetChallenge(), req.GetValidate())
	if err != nil {
		return nil, err
	}
	return &pb.NewSMSResp{
		CaptchaKey: captchaKey,
	}, nil
}

func (s *BilibiliService) LoginWithSMS(ctx context.Context, req *pb.LoginWithSMSReq) (*pb.LoginWithSMSResp, error) {
	c, err := bilibili.LoginWithSMS(ctx, req.GetPhone(), req.GetCode(), req.GetCaptchaKey())
	if err != nil {
		return nil, err
	}
	return &pb.LoginWithSMSResp{
		Cookies: map[string]string{c.Name: c.Value},
	}, nil
}

func (s *BilibiliService) ParseVideoPage(ctx context.Context, req *pb.ParseVideoPageReq) (*pb.VideoPageInfo, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.GetCookies()), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	r, err := c.ParseVideoPage(req.GetAid(), req.GetBvid())
	if err != nil {
		return nil, err
	}
	resp := &pb.VideoPageInfo{
		Title:      r.Title,
		Actors:     r.Actors,
		VideoInfos: videoInfos2pb(r.VideoInfos),
	}
	return resp, nil
}

func (s *BilibiliService) GetVideoURL(ctx context.Context, req *pb.GetVideoURLReq) (*pb.VideoURL, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.GetCookies()), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	conf := []bilibili.GetVideoURLConfig{}
	if req.GetQuality() != 0 {
		conf = append(conf, bilibili.WithQuality(req.GetQuality()))
	}
	r, err := c.GetVideoURL(req.GetAid(), req.GetBvid(), req.GetCid(), conf...)
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
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.GetCookies()), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	opt := make([]bilibili.GetDashVideoURLConfig, 0)
	m, m2, err := c.GetDashVideoURL(req.GetAid(), req.GetBvid(), req.GetCid(), opt...)
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
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.GetCookies()), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	sub, err := c.GetSubtitles(req.GetAid(), req.GetBvid(), req.GetCid())
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
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.GetCookies()), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	r, err := c.ParsePGCPage(req.GetEpid(), req.GetSsid())
	if err != nil {
		return nil, err
	}
	resp := &pb.VideoPageInfo{
		Title:      r.Title,
		Actors:     r.Actors,
		VideoInfos: videoInfos2pb(r.VideoInfos),
	}
	return resp, nil
}

func (s *BilibiliService) GetPGCURL(ctx context.Context, req *pb.GetPGCURLReq) (*pb.VideoURL, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.GetCookies()), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	conf := []bilibili.GetVideoURLConfig{}
	if req.GetQuality() != 0 {
		conf = append(conf, bilibili.WithQuality(req.GetQuality()))
	}
	r, err := c.GetPGCURL(req.GetEpid(), req.GetCid(), conf...)
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
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.GetCookies()), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	opt := make([]bilibili.GetDashVideoURLConfig, 0)
	m, m2, err := c.GetDashPGCURL(req.GetEpid(), req.GetCid(), opt...)
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
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.GetCookies()), bilibili.WithContext(ctx))
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
	t, id, err := bilibili.Match(req.GetUrl())
	if err != nil {
		return nil, err
	}
	return &pb.MatchResp{
		Type: t,
		Id:   id,
	}, nil
}

func (s *BilibiliService) GetLiveStreams(ctx context.Context, req *pb.GetLiveStreamsReq) (*pb.GetLiveStreamsResp, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.GetCookies()), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	r, err := c.GetLiveStreams(req.GetCid(), req.GetHls())
	if err != nil {
		return nil, err
	}
	resp := &pb.GetLiveStreamsResp{
		LiveStreams: make([]*pb.LiveStream, len(r)),
	}
	for i, v := range r {
		resp.LiveStreams[i] = &pb.LiveStream{
			Quality: v.Quality,
			Urls:    v.Urls,
			Desc:    v.Desc,
		}
	}
	return resp, nil
}

func (s *BilibiliService) ParseLivePage(ctx context.Context, req *pb.ParseLivePageReq) (*pb.VideoPageInfo, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.GetCookies()), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	r, err := c.ParseLivePage(req.GetRoomID())
	if err != nil {
		return nil, err
	}
	resp := &pb.VideoPageInfo{
		Title:      r.Title,
		Actors:     r.Actors,
		VideoInfos: videoInfos2pb(r.VideoInfos),
	}
	return resp, nil
}

func videoInfos2pb(in []*bilibili.VideoInfo) []*pb.VideoInfo {
	if len(in) == 0 {
		return nil
	}
	out := make([]*pb.VideoInfo, len(in))
	for i, v := range in {
		out[i] = &pb.VideoInfo{
			Bvid:       v.Bvid,
			Cid:        v.Cid,
			Epid:       v.Epid,
			Name:       v.Name,
			CoverImage: v.CoverImage,
			Live:       v.Live,
		}
	}
	return out
}

func (s *BilibiliService) GetLiveDanmuInfo(ctx context.Context, req *pb.GetLiveDanmuInfoReq) (*pb.GetLiveDanmuInfoResp, error) {
	c, err := bilibili.NewClient(utils.MapToHttpCookies(req.GetCookies()), bilibili.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	r, err := c.GetLiveDanmuInfo(req.GetRoomID())
	if err != nil {
		return nil, err
	}
	resp := &pb.GetLiveDanmuInfoResp{
		Token: r.Token,
	}
	resp.HostList = make([]*pb.GetLiveDanmuInfoResp_Host, len(r.HostList))
	for i, v := range r.HostList {
		resp.HostList[i] = &pb.GetLiveDanmuInfoResp_Host{
			Host:    v.Host,
			Port:    uint32(v.Port),
			WssPort: uint32(v.WssPort),
			WsPort:  uint32(v.WsPort),
		}
	}
	return resp, nil
}
