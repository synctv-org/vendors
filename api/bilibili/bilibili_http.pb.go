// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.0
// source: bilibili/bilibili.proto

package bilibili

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBilibiliGetDashPGCURL = "/api.bilibili.Bilibili/GetDashPGCURL"
const OperationBilibiliGetDashVideoURL = "/api.bilibili.Bilibili/GetDashVideoURL"
const OperationBilibiliGetPGCURL = "/api.bilibili.Bilibili/GetPGCURL"
const OperationBilibiliGetSubtitles = "/api.bilibili.Bilibili/GetSubtitles"
const OperationBilibiliGetVideoURL = "/api.bilibili.Bilibili/GetVideoURL"
const OperationBilibiliLoginWithQRCode = "/api.bilibili.Bilibili/LoginWithQRCode"
const OperationBilibiliLoginWithSMS = "/api.bilibili.Bilibili/LoginWithSMS"
const OperationBilibiliMatch = "/api.bilibili.Bilibili/Match"
const OperationBilibiliNewCaptcha = "/api.bilibili.Bilibili/NewCaptcha"
const OperationBilibiliNewQRCode = "/api.bilibili.Bilibili/NewQRCode"
const OperationBilibiliNewSMS = "/api.bilibili.Bilibili/NewSMS"
const OperationBilibiliParsePGCPage = "/api.bilibili.Bilibili/ParsePGCPage"
const OperationBilibiliParseVideoPage = "/api.bilibili.Bilibili/ParseVideoPage"
const OperationBilibiliUserInfo = "/api.bilibili.Bilibili/UserInfo"

type BilibiliHTTPServer interface {
	GetDashPGCURL(context.Context, *GetDashPGCURLReq) (*GetDashPGCURLResp, error)
	GetDashVideoURL(context.Context, *GetDashVideoURLReq) (*GetDashVideoURLResp, error)
	GetPGCURL(context.Context, *GetPGCURLReq) (*VideoURL, error)
	GetSubtitles(context.Context, *GetSubtitlesReq) (*GetSubtitlesResp, error)
	GetVideoURL(context.Context, *GetVideoURLReq) (*VideoURL, error)
	LoginWithQRCode(context.Context, *LoginWithQRCodeReq) (*LoginWithQRCodeResp, error)
	LoginWithSMS(context.Context, *LoginWithSMSReq) (*LoginWithSMSResp, error)
	Match(context.Context, *MatchReq) (*MatchResp, error)
	NewCaptcha(context.Context, *Empty) (*NewCaptchaResp, error)
	NewQRCode(context.Context, *Empty) (*NewQRCodeResp, error)
	NewSMS(context.Context, *NewSMSReq) (*NewSMSResp, error)
	ParsePGCPage(context.Context, *ParsePGCPageReq) (*VideoPageInfo, error)
	ParseVideoPage(context.Context, *ParseVideoPageReq) (*VideoPageInfo, error)
	UserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error)
}

func RegisterBilibiliHTTPServer(s *http.Server, srv BilibiliHTTPServer) {
	r := s.Route("/")
	r.GET("/api/bilibili/new_qrcode", _Bilibili_NewQRCode0_HTTP_Handler(srv))
	r.POST("/api/bilibili/login_with_qrcode", _Bilibili_LoginWithQRCode0_HTTP_Handler(srv))
	r.GET("/api/bilibili/new_captcha", _Bilibili_NewCaptcha0_HTTP_Handler(srv))
	r.POST("/api/bilibili/new_sms", _Bilibili_NewSMS0_HTTP_Handler(srv))
	r.POST("/api/bilibili/login_with_sms", _Bilibili_LoginWithSMS0_HTTP_Handler(srv))
	r.POST("/api/bilibili/parse_video_page", _Bilibili_ParseVideoPage0_HTTP_Handler(srv))
	r.POST("/api/bilibili/get_video_url", _Bilibili_GetVideoURL0_HTTP_Handler(srv))
	r.POST("/api/bilibili/get_dash_video_url", _Bilibili_GetDashVideoURL0_HTTP_Handler(srv))
	r.POST("/api/bilibili/get_subtitles", _Bilibili_GetSubtitles0_HTTP_Handler(srv))
	r.POST("/api/bilibili/parse_pgc_page", _Bilibili_ParsePGCPage0_HTTP_Handler(srv))
	r.POST("/api/bilibili/get_pgc_url", _Bilibili_GetPGCURL0_HTTP_Handler(srv))
	r.POST("/api/bilibili/get_dash_pgc_url", _Bilibili_GetDashPGCURL0_HTTP_Handler(srv))
	r.POST("/api/bilibili/user_info", _Bilibili_UserInfo0_HTTP_Handler(srv))
	r.POST("/api/bilibili/match", _Bilibili_Match0_HTTP_Handler(srv))
}

func _Bilibili_NewQRCode0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliNewQRCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NewQRCode(ctx, req.(*Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NewQRCodeResp)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_LoginWithQRCode0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginWithQRCodeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliLoginWithQRCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginWithQRCode(ctx, req.(*LoginWithQRCodeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginWithQRCodeResp)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_NewCaptcha0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliNewCaptcha)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NewCaptcha(ctx, req.(*Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NewCaptchaResp)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_NewSMS0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NewSMSReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliNewSMS)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NewSMS(ctx, req.(*NewSMSReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NewSMSResp)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_LoginWithSMS0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginWithSMSReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliLoginWithSMS)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginWithSMS(ctx, req.(*LoginWithSMSReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginWithSMSResp)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_ParseVideoPage0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ParseVideoPageReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliParseVideoPage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ParseVideoPage(ctx, req.(*ParseVideoPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VideoPageInfo)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_GetVideoURL0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVideoURLReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliGetVideoURL)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVideoURL(ctx, req.(*GetVideoURLReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VideoURL)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_GetDashVideoURL0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDashVideoURLReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliGetDashVideoURL)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDashVideoURL(ctx, req.(*GetDashVideoURLReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDashVideoURLResp)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_GetSubtitles0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSubtitlesReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliGetSubtitles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSubtitles(ctx, req.(*GetSubtitlesReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSubtitlesResp)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_ParsePGCPage0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ParsePGCPageReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliParsePGCPage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ParsePGCPage(ctx, req.(*ParsePGCPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VideoPageInfo)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_GetPGCURL0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPGCURLReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliGetPGCURL)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPGCURL(ctx, req.(*GetPGCURLReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VideoURL)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_GetDashPGCURL0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDashPGCURLReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliGetDashPGCURL)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDashPGCURL(ctx, req.(*GetDashPGCURLReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDashPGCURLResp)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_UserInfo0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfoReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserInfo(ctx, req.(*UserInfoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoResp)
		return ctx.Result(200, reply)
	}
}

func _Bilibili_Match0_HTTP_Handler(srv BilibiliHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MatchReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBilibiliMatch)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Match(ctx, req.(*MatchReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MatchResp)
		return ctx.Result(200, reply)
	}
}

type BilibiliHTTPClient interface {
	GetDashPGCURL(ctx context.Context, req *GetDashPGCURLReq, opts ...http.CallOption) (rsp *GetDashPGCURLResp, err error)
	GetDashVideoURL(ctx context.Context, req *GetDashVideoURLReq, opts ...http.CallOption) (rsp *GetDashVideoURLResp, err error)
	GetPGCURL(ctx context.Context, req *GetPGCURLReq, opts ...http.CallOption) (rsp *VideoURL, err error)
	GetSubtitles(ctx context.Context, req *GetSubtitlesReq, opts ...http.CallOption) (rsp *GetSubtitlesResp, err error)
	GetVideoURL(ctx context.Context, req *GetVideoURLReq, opts ...http.CallOption) (rsp *VideoURL, err error)
	LoginWithQRCode(ctx context.Context, req *LoginWithQRCodeReq, opts ...http.CallOption) (rsp *LoginWithQRCodeResp, err error)
	LoginWithSMS(ctx context.Context, req *LoginWithSMSReq, opts ...http.CallOption) (rsp *LoginWithSMSResp, err error)
	Match(ctx context.Context, req *MatchReq, opts ...http.CallOption) (rsp *MatchResp, err error)
	NewCaptcha(ctx context.Context, req *Empty, opts ...http.CallOption) (rsp *NewCaptchaResp, err error)
	NewQRCode(ctx context.Context, req *Empty, opts ...http.CallOption) (rsp *NewQRCodeResp, err error)
	NewSMS(ctx context.Context, req *NewSMSReq, opts ...http.CallOption) (rsp *NewSMSResp, err error)
	ParsePGCPage(ctx context.Context, req *ParsePGCPageReq, opts ...http.CallOption) (rsp *VideoPageInfo, err error)
	ParseVideoPage(ctx context.Context, req *ParseVideoPageReq, opts ...http.CallOption) (rsp *VideoPageInfo, err error)
	UserInfo(ctx context.Context, req *UserInfoReq, opts ...http.CallOption) (rsp *UserInfoResp, err error)
}

type BilibiliHTTPClientImpl struct {
	cc *http.Client
}

func NewBilibiliHTTPClient(client *http.Client) BilibiliHTTPClient {
	return &BilibiliHTTPClientImpl{client}
}

func (c *BilibiliHTTPClientImpl) GetDashPGCURL(ctx context.Context, in *GetDashPGCURLReq, opts ...http.CallOption) (*GetDashPGCURLResp, error) {
	var out GetDashPGCURLResp
	pattern := "/api/bilibili/get_dash_pgc_url"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliGetDashPGCURL))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) GetDashVideoURL(ctx context.Context, in *GetDashVideoURLReq, opts ...http.CallOption) (*GetDashVideoURLResp, error) {
	var out GetDashVideoURLResp
	pattern := "/api/bilibili/get_dash_video_url"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliGetDashVideoURL))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) GetPGCURL(ctx context.Context, in *GetPGCURLReq, opts ...http.CallOption) (*VideoURL, error) {
	var out VideoURL
	pattern := "/api/bilibili/get_pgc_url"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliGetPGCURL))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) GetSubtitles(ctx context.Context, in *GetSubtitlesReq, opts ...http.CallOption) (*GetSubtitlesResp, error) {
	var out GetSubtitlesResp
	pattern := "/api/bilibili/get_subtitles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliGetSubtitles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) GetVideoURL(ctx context.Context, in *GetVideoURLReq, opts ...http.CallOption) (*VideoURL, error) {
	var out VideoURL
	pattern := "/api/bilibili/get_video_url"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliGetVideoURL))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) LoginWithQRCode(ctx context.Context, in *LoginWithQRCodeReq, opts ...http.CallOption) (*LoginWithQRCodeResp, error) {
	var out LoginWithQRCodeResp
	pattern := "/api/bilibili/login_with_qrcode"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliLoginWithQRCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) LoginWithSMS(ctx context.Context, in *LoginWithSMSReq, opts ...http.CallOption) (*LoginWithSMSResp, error) {
	var out LoginWithSMSResp
	pattern := "/api/bilibili/login_with_sms"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliLoginWithSMS))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) Match(ctx context.Context, in *MatchReq, opts ...http.CallOption) (*MatchResp, error) {
	var out MatchResp
	pattern := "/api/bilibili/match"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliMatch))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) NewCaptcha(ctx context.Context, in *Empty, opts ...http.CallOption) (*NewCaptchaResp, error) {
	var out NewCaptchaResp
	pattern := "/api/bilibili/new_captcha"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBilibiliNewCaptcha))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) NewQRCode(ctx context.Context, in *Empty, opts ...http.CallOption) (*NewQRCodeResp, error) {
	var out NewQRCodeResp
	pattern := "/api/bilibili/new_qrcode"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBilibiliNewQRCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) NewSMS(ctx context.Context, in *NewSMSReq, opts ...http.CallOption) (*NewSMSResp, error) {
	var out NewSMSResp
	pattern := "/api/bilibili/new_sms"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliNewSMS))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) ParsePGCPage(ctx context.Context, in *ParsePGCPageReq, opts ...http.CallOption) (*VideoPageInfo, error) {
	var out VideoPageInfo
	pattern := "/api/bilibili/parse_pgc_page"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliParsePGCPage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) ParseVideoPage(ctx context.Context, in *ParseVideoPageReq, opts ...http.CallOption) (*VideoPageInfo, error) {
	var out VideoPageInfo
	pattern := "/api/bilibili/parse_video_page"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliParseVideoPage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BilibiliHTTPClientImpl) UserInfo(ctx context.Context, in *UserInfoReq, opts ...http.CallOption) (*UserInfoResp, error) {
	var out UserInfoResp
	pattern := "/api/bilibili/user_info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBilibiliUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
