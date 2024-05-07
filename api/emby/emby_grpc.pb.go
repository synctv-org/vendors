// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: emby/emby.proto

package emby

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
	Emby_Login_FullMethodName                  = "/api.emby.Emby/Login"
	Emby_Me_FullMethodName                     = "/api.emby.Emby/Me"
	Emby_GetItems_FullMethodName               = "/api.emby.Emby/GetItems"
	Emby_GetItem_FullMethodName                = "/api.emby.Emby/GetItem"
	Emby_GetSystemInfo_FullMethodName          = "/api.emby.Emby/GetSystemInfo"
	Emby_FsList_FullMethodName                 = "/api.emby.Emby/FsList"
	Emby_Logout_FullMethodName                 = "/api.emby.Emby/Logout"
	Emby_PlaybackInfo_FullMethodName           = "/api.emby.Emby/PlaybackInfo"
	Emby_DeleteActiveEncodeings_FullMethodName = "/api.emby.Emby/DeleteActiveEncodeings"
)

// EmbyClient is the client API for Emby service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmbyClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Me(ctx context.Context, in *MeReq, opts ...grpc.CallOption) (*MeResp, error)
	GetItems(ctx context.Context, in *GetItemsReq, opts ...grpc.CallOption) (*GetItemsResp, error)
	GetItem(ctx context.Context, in *GetItemReq, opts ...grpc.CallOption) (*Item, error)
	GetSystemInfo(ctx context.Context, in *SystemInfoReq, opts ...grpc.CallOption) (*SystemInfoResp, error)
	FsList(ctx context.Context, in *FsListReq, opts ...grpc.CallOption) (*FsListResp, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*Empty, error)
	PlaybackInfo(ctx context.Context, in *PlaybackInfoReq, opts ...grpc.CallOption) (*PlaybackInfoResp, error)
	DeleteActiveEncodeings(ctx context.Context, in *DeleteActiveEncodeingsReq, opts ...grpc.CallOption) (*Empty, error)
}

type embyClient struct {
	cc grpc.ClientConnInterface
}

func NewEmbyClient(cc grpc.ClientConnInterface) EmbyClient {
	return &embyClient{cc}
}

func (c *embyClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Emby_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *embyClient) Me(ctx context.Context, in *MeReq, opts ...grpc.CallOption) (*MeResp, error) {
	out := new(MeResp)
	err := c.cc.Invoke(ctx, Emby_Me_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *embyClient) GetItems(ctx context.Context, in *GetItemsReq, opts ...grpc.CallOption) (*GetItemsResp, error) {
	out := new(GetItemsResp)
	err := c.cc.Invoke(ctx, Emby_GetItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *embyClient) GetItem(ctx context.Context, in *GetItemReq, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, Emby_GetItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *embyClient) GetSystemInfo(ctx context.Context, in *SystemInfoReq, opts ...grpc.CallOption) (*SystemInfoResp, error) {
	out := new(SystemInfoResp)
	err := c.cc.Invoke(ctx, Emby_GetSystemInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *embyClient) FsList(ctx context.Context, in *FsListReq, opts ...grpc.CallOption) (*FsListResp, error) {
	out := new(FsListResp)
	err := c.cc.Invoke(ctx, Emby_FsList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *embyClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Emby_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *embyClient) PlaybackInfo(ctx context.Context, in *PlaybackInfoReq, opts ...grpc.CallOption) (*PlaybackInfoResp, error) {
	out := new(PlaybackInfoResp)
	err := c.cc.Invoke(ctx, Emby_PlaybackInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *embyClient) DeleteActiveEncodeings(ctx context.Context, in *DeleteActiveEncodeingsReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Emby_DeleteActiveEncodeings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmbyServer is the server API for Emby service.
// All implementations must embed UnimplementedEmbyServer
// for forward compatibility
type EmbyServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Me(context.Context, *MeReq) (*MeResp, error)
	GetItems(context.Context, *GetItemsReq) (*GetItemsResp, error)
	GetItem(context.Context, *GetItemReq) (*Item, error)
	GetSystemInfo(context.Context, *SystemInfoReq) (*SystemInfoResp, error)
	FsList(context.Context, *FsListReq) (*FsListResp, error)
	Logout(context.Context, *LogoutReq) (*Empty, error)
	PlaybackInfo(context.Context, *PlaybackInfoReq) (*PlaybackInfoResp, error)
	DeleteActiveEncodeings(context.Context, *DeleteActiveEncodeingsReq) (*Empty, error)
	mustEmbedUnimplementedEmbyServer()
}

// UnimplementedEmbyServer must be embedded to have forward compatible implementations.
type UnimplementedEmbyServer struct {
}

func (UnimplementedEmbyServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedEmbyServer) Me(context.Context, *MeReq) (*MeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Me not implemented")
}
func (UnimplementedEmbyServer) GetItems(context.Context, *GetItemsReq) (*GetItemsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}
func (UnimplementedEmbyServer) GetItem(context.Context, *GetItemReq) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedEmbyServer) GetSystemInfo(context.Context, *SystemInfoReq) (*SystemInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemInfo not implemented")
}
func (UnimplementedEmbyServer) FsList(context.Context, *FsListReq) (*FsListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FsList not implemented")
}
func (UnimplementedEmbyServer) Logout(context.Context, *LogoutReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedEmbyServer) PlaybackInfo(context.Context, *PlaybackInfoReq) (*PlaybackInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaybackInfo not implemented")
}
func (UnimplementedEmbyServer) DeleteActiveEncodeings(context.Context, *DeleteActiveEncodeingsReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteActiveEncodeings not implemented")
}
func (UnimplementedEmbyServer) mustEmbedUnimplementedEmbyServer() {}

// UnsafeEmbyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmbyServer will
// result in compilation errors.
type UnsafeEmbyServer interface {
	mustEmbedUnimplementedEmbyServer()
}

func RegisterEmbyServer(s grpc.ServiceRegistrar, srv EmbyServer) {
	s.RegisterService(&Emby_ServiceDesc, srv)
}

func _Emby_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbyServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Emby_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbyServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Emby_Me_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbyServer).Me(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Emby_Me_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbyServer).Me(ctx, req.(*MeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Emby_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbyServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Emby_GetItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbyServer).GetItems(ctx, req.(*GetItemsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Emby_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbyServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Emby_GetItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbyServer).GetItem(ctx, req.(*GetItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Emby_GetSystemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbyServer).GetSystemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Emby_GetSystemInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbyServer).GetSystemInfo(ctx, req.(*SystemInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Emby_FsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FsListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbyServer).FsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Emby_FsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbyServer).FsList(ctx, req.(*FsListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Emby_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbyServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Emby_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbyServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Emby_PlaybackInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaybackInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbyServer).PlaybackInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Emby_PlaybackInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbyServer).PlaybackInfo(ctx, req.(*PlaybackInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Emby_DeleteActiveEncodeings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActiveEncodeingsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbyServer).DeleteActiveEncodeings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Emby_DeleteActiveEncodeings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbyServer).DeleteActiveEncodeings(ctx, req.(*DeleteActiveEncodeingsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Emby_ServiceDesc is the grpc.ServiceDesc for Emby service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Emby_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.emby.Emby",
	HandlerType: (*EmbyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Emby_Login_Handler,
		},
		{
			MethodName: "Me",
			Handler:    _Emby_Me_Handler,
		},
		{
			MethodName: "GetItems",
			Handler:    _Emby_GetItems_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _Emby_GetItem_Handler,
		},
		{
			MethodName: "GetSystemInfo",
			Handler:    _Emby_GetSystemInfo_Handler,
		},
		{
			MethodName: "FsList",
			Handler:    _Emby_FsList_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Emby_Logout_Handler,
		},
		{
			MethodName: "PlaybackInfo",
			Handler:    _Emby_PlaybackInfo_Handler,
		},
		{
			MethodName: "DeleteActiveEncodeings",
			Handler:    _Emby_DeleteActiveEncodeings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emby/emby.proto",
}
