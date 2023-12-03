// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bilibili/bilibili.proto

package bilibili

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Bilibili_NewQRCode_FullMethodName       = "/api.bilibili.Bilibili/NewQRCode"
	Bilibili_LoginWithQRCode_FullMethodName = "/api.bilibili.Bilibili/LoginWithQRCode"
	Bilibili_NewCaptcha_FullMethodName      = "/api.bilibili.Bilibili/NewCaptcha"
	Bilibili_NewSMS_FullMethodName          = "/api.bilibili.Bilibili/NewSMS"
	Bilibili_LoginWithSMS_FullMethodName    = "/api.bilibili.Bilibili/LoginWithSMS"
	Bilibili_ParseVideoPage_FullMethodName  = "/api.bilibili.Bilibili/ParseVideoPage"
	Bilibili_GetVideoURL_FullMethodName     = "/api.bilibili.Bilibili/GetVideoURL"
	Bilibili_GetDashVideoURL_FullMethodName = "/api.bilibili.Bilibili/GetDashVideoURL"
	Bilibili_GetSubtitles_FullMethodName    = "/api.bilibili.Bilibili/GetSubtitles"
	Bilibili_ParsePGCPage_FullMethodName    = "/api.bilibili.Bilibili/ParsePGCPage"
	Bilibili_GetPGCURL_FullMethodName       = "/api.bilibili.Bilibili/GetPGCURL"
	Bilibili_GetDashPGCURL_FullMethodName   = "/api.bilibili.Bilibili/GetDashPGCURL"
	Bilibili_UserInfo_FullMethodName        = "/api.bilibili.Bilibili/UserInfo"
	Bilibili_Match_FullMethodName           = "/api.bilibili.Bilibili/Match"
)

// BilibiliClient is the client API for Bilibili service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BilibiliClient interface {
	NewQRCode(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NewQRCodeResp, error)
	LoginWithQRCode(ctx context.Context, in *LoginWithQRCodeReq, opts ...grpc.CallOption) (*LoginWithQRCodeResp, error)
	NewCaptcha(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NewCaptchaResp, error)
	NewSMS(ctx context.Context, in *NewSMSReq, opts ...grpc.CallOption) (*NewSMSResp, error)
	LoginWithSMS(ctx context.Context, in *LoginWithSMSReq, opts ...grpc.CallOption) (*LoginWithSMSResp, error)
	ParseVideoPage(ctx context.Context, in *ParseVideoPageReq, opts ...grpc.CallOption) (*VideoPageInfo, error)
	GetVideoURL(ctx context.Context, in *GetVideoURLReq, opts ...grpc.CallOption) (*VideoURL, error)
	GetDashVideoURL(ctx context.Context, in *GetDashVideoURLReq, opts ...grpc.CallOption) (*GetDashVideoURLResp, error)
	GetSubtitles(ctx context.Context, in *GetSubtitlesReq, opts ...grpc.CallOption) (*GetSubtitlesResp, error)
	ParsePGCPage(ctx context.Context, in *ParsePGCPageReq, opts ...grpc.CallOption) (*VideoPageInfo, error)
	GetPGCURL(ctx context.Context, in *GetPGCURLReq, opts ...grpc.CallOption) (*VideoURL, error)
	GetDashPGCURL(ctx context.Context, in *GetDashPGCURLReq, opts ...grpc.CallOption) (*GetDashPGCURLResp, error)
	UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
	Match(ctx context.Context, in *MatchReq, opts ...grpc.CallOption) (*MatchResp, error)
}

type bilibiliClient struct {
	cc grpc.ClientConnInterface
}

func NewBilibiliClient(cc grpc.ClientConnInterface) BilibiliClient {
	return &bilibiliClient{cc}
}

func (c *bilibiliClient) NewQRCode(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NewQRCodeResp, error) {
	out := new(NewQRCodeResp)
	err := c.cc.Invoke(ctx, Bilibili_NewQRCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) LoginWithQRCode(ctx context.Context, in *LoginWithQRCodeReq, opts ...grpc.CallOption) (*LoginWithQRCodeResp, error) {
	out := new(LoginWithQRCodeResp)
	err := c.cc.Invoke(ctx, Bilibili_LoginWithQRCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) NewCaptcha(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NewCaptchaResp, error) {
	out := new(NewCaptchaResp)
	err := c.cc.Invoke(ctx, Bilibili_NewCaptcha_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) NewSMS(ctx context.Context, in *NewSMSReq, opts ...grpc.CallOption) (*NewSMSResp, error) {
	out := new(NewSMSResp)
	err := c.cc.Invoke(ctx, Bilibili_NewSMS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) LoginWithSMS(ctx context.Context, in *LoginWithSMSReq, opts ...grpc.CallOption) (*LoginWithSMSResp, error) {
	out := new(LoginWithSMSResp)
	err := c.cc.Invoke(ctx, Bilibili_LoginWithSMS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) ParseVideoPage(ctx context.Context, in *ParseVideoPageReq, opts ...grpc.CallOption) (*VideoPageInfo, error) {
	out := new(VideoPageInfo)
	err := c.cc.Invoke(ctx, Bilibili_ParseVideoPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) GetVideoURL(ctx context.Context, in *GetVideoURLReq, opts ...grpc.CallOption) (*VideoURL, error) {
	out := new(VideoURL)
	err := c.cc.Invoke(ctx, Bilibili_GetVideoURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) GetDashVideoURL(ctx context.Context, in *GetDashVideoURLReq, opts ...grpc.CallOption) (*GetDashVideoURLResp, error) {
	out := new(GetDashVideoURLResp)
	err := c.cc.Invoke(ctx, Bilibili_GetDashVideoURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) GetSubtitles(ctx context.Context, in *GetSubtitlesReq, opts ...grpc.CallOption) (*GetSubtitlesResp, error) {
	out := new(GetSubtitlesResp)
	err := c.cc.Invoke(ctx, Bilibili_GetSubtitles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) ParsePGCPage(ctx context.Context, in *ParsePGCPageReq, opts ...grpc.CallOption) (*VideoPageInfo, error) {
	out := new(VideoPageInfo)
	err := c.cc.Invoke(ctx, Bilibili_ParsePGCPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) GetPGCURL(ctx context.Context, in *GetPGCURLReq, opts ...grpc.CallOption) (*VideoURL, error) {
	out := new(VideoURL)
	err := c.cc.Invoke(ctx, Bilibili_GetPGCURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) GetDashPGCURL(ctx context.Context, in *GetDashPGCURLReq, opts ...grpc.CallOption) (*GetDashPGCURLResp, error) {
	out := new(GetDashPGCURLResp)
	err := c.cc.Invoke(ctx, Bilibili_GetDashPGCURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, Bilibili_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bilibiliClient) Match(ctx context.Context, in *MatchReq, opts ...grpc.CallOption) (*MatchResp, error) {
	out := new(MatchResp)
	err := c.cc.Invoke(ctx, Bilibili_Match_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BilibiliServer is the server API for Bilibili service.
// All implementations must embed UnimplementedBilibiliServer
// for forward compatibility
type BilibiliServer interface {
	NewQRCode(context.Context, *Empty) (*NewQRCodeResp, error)
	LoginWithQRCode(context.Context, *LoginWithQRCodeReq) (*LoginWithQRCodeResp, error)
	NewCaptcha(context.Context, *Empty) (*NewCaptchaResp, error)
	NewSMS(context.Context, *NewSMSReq) (*NewSMSResp, error)
	LoginWithSMS(context.Context, *LoginWithSMSReq) (*LoginWithSMSResp, error)
	ParseVideoPage(context.Context, *ParseVideoPageReq) (*VideoPageInfo, error)
	GetVideoURL(context.Context, *GetVideoURLReq) (*VideoURL, error)
	GetDashVideoURL(context.Context, *GetDashVideoURLReq) (*GetDashVideoURLResp, error)
	GetSubtitles(context.Context, *GetSubtitlesReq) (*GetSubtitlesResp, error)
	ParsePGCPage(context.Context, *ParsePGCPageReq) (*VideoPageInfo, error)
	GetPGCURL(context.Context, *GetPGCURLReq) (*VideoURL, error)
	GetDashPGCURL(context.Context, *GetDashPGCURLReq) (*GetDashPGCURLResp, error)
	UserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error)
	Match(context.Context, *MatchReq) (*MatchResp, error)
	mustEmbedUnimplementedBilibiliServer()
}

// UnimplementedBilibiliServer must be embedded to have forward compatible implementations.
type UnimplementedBilibiliServer struct {
}

func (UnimplementedBilibiliServer) NewQRCode(context.Context, *Empty) (*NewQRCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewQRCode not implemented")
}
func (UnimplementedBilibiliServer) LoginWithQRCode(context.Context, *LoginWithQRCodeReq) (*LoginWithQRCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithQRCode not implemented")
}
func (UnimplementedBilibiliServer) NewCaptcha(context.Context, *Empty) (*NewCaptchaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewCaptcha not implemented")
}
func (UnimplementedBilibiliServer) NewSMS(context.Context, *NewSMSReq) (*NewSMSResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSMS not implemented")
}
func (UnimplementedBilibiliServer) LoginWithSMS(context.Context, *LoginWithSMSReq) (*LoginWithSMSResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithSMS not implemented")
}
func (UnimplementedBilibiliServer) ParseVideoPage(context.Context, *ParseVideoPageReq) (*VideoPageInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseVideoPage not implemented")
}
func (UnimplementedBilibiliServer) GetVideoURL(context.Context, *GetVideoURLReq) (*VideoURL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoURL not implemented")
}
func (UnimplementedBilibiliServer) GetDashVideoURL(context.Context, *GetDashVideoURLReq) (*GetDashVideoURLResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDashVideoURL not implemented")
}
func (UnimplementedBilibiliServer) GetSubtitles(context.Context, *GetSubtitlesReq) (*GetSubtitlesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubtitles not implemented")
}
func (UnimplementedBilibiliServer) ParsePGCPage(context.Context, *ParsePGCPageReq) (*VideoPageInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParsePGCPage not implemented")
}
func (UnimplementedBilibiliServer) GetPGCURL(context.Context, *GetPGCURLReq) (*VideoURL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPGCURL not implemented")
}
func (UnimplementedBilibiliServer) GetDashPGCURL(context.Context, *GetDashPGCURLReq) (*GetDashPGCURLResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDashPGCURL not implemented")
}
func (UnimplementedBilibiliServer) UserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedBilibiliServer) Match(context.Context, *MatchReq) (*MatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Match not implemented")
}
func (UnimplementedBilibiliServer) mustEmbedUnimplementedBilibiliServer() {}

// UnsafeBilibiliServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BilibiliServer will
// result in compilation errors.
type UnsafeBilibiliServer interface {
	mustEmbedUnimplementedBilibiliServer()
}

func RegisterBilibiliServer(s grpc.ServiceRegistrar, srv BilibiliServer) {
	s.RegisterService(&Bilibili_ServiceDesc, srv)
}

func _Bilibili_NewQRCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).NewQRCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_NewQRCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).NewQRCode(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_LoginWithQRCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithQRCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).LoginWithQRCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_LoginWithQRCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).LoginWithQRCode(ctx, req.(*LoginWithQRCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_NewCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).NewCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_NewCaptcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).NewCaptcha(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_NewSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewSMSReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).NewSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_NewSMS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).NewSMS(ctx, req.(*NewSMSReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_LoginWithSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithSMSReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).LoginWithSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_LoginWithSMS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).LoginWithSMS(ctx, req.(*LoginWithSMSReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_ParseVideoPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseVideoPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).ParseVideoPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_ParseVideoPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).ParseVideoPage(ctx, req.(*ParseVideoPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_GetVideoURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).GetVideoURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_GetVideoURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).GetVideoURL(ctx, req.(*GetVideoURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_GetDashVideoURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDashVideoURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).GetDashVideoURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_GetDashVideoURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).GetDashVideoURL(ctx, req.(*GetDashVideoURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_GetSubtitles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubtitlesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).GetSubtitles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_GetSubtitles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).GetSubtitles(ctx, req.(*GetSubtitlesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_ParsePGCPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParsePGCPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).ParsePGCPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_ParsePGCPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).ParsePGCPage(ctx, req.(*ParsePGCPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_GetPGCURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPGCURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).GetPGCURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_GetPGCURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).GetPGCURL(ctx, req.(*GetPGCURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_GetDashPGCURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDashPGCURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).GetDashPGCURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_GetDashPGCURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).GetDashPGCURL(ctx, req.(*GetDashPGCURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).UserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bilibili_Match_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilibiliServer).Match(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bilibili_Match_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilibiliServer).Match(ctx, req.(*MatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Bilibili_ServiceDesc is the grpc.ServiceDesc for Bilibili service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bilibili_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.bilibili.Bilibili",
	HandlerType: (*BilibiliServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewQRCode",
			Handler:    _Bilibili_NewQRCode_Handler,
		},
		{
			MethodName: "LoginWithQRCode",
			Handler:    _Bilibili_LoginWithQRCode_Handler,
		},
		{
			MethodName: "NewCaptcha",
			Handler:    _Bilibili_NewCaptcha_Handler,
		},
		{
			MethodName: "NewSMS",
			Handler:    _Bilibili_NewSMS_Handler,
		},
		{
			MethodName: "LoginWithSMS",
			Handler:    _Bilibili_LoginWithSMS_Handler,
		},
		{
			MethodName: "ParseVideoPage",
			Handler:    _Bilibili_ParseVideoPage_Handler,
		},
		{
			MethodName: "GetVideoURL",
			Handler:    _Bilibili_GetVideoURL_Handler,
		},
		{
			MethodName: "GetDashVideoURL",
			Handler:    _Bilibili_GetDashVideoURL_Handler,
		},
		{
			MethodName: "GetSubtitles",
			Handler:    _Bilibili_GetSubtitles_Handler,
		},
		{
			MethodName: "ParsePGCPage",
			Handler:    _Bilibili_ParsePGCPage_Handler,
		},
		{
			MethodName: "GetPGCURL",
			Handler:    _Bilibili_GetPGCURL_Handler,
		},
		{
			MethodName: "GetDashPGCURL",
			Handler:    _Bilibili_GetDashPGCURL_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _Bilibili_UserInfo_Handler,
		},
		{
			MethodName: "Match",
			Handler:    _Bilibili_Match_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/bilibili.proto",
}
