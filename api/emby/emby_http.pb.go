// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.3
// - protoc             v5.29.3
// source: emby/emby.proto

package emby

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

const OperationEmbyDeleteActiveEncodeings = "/api.emby.Emby/DeleteActiveEncodeings"
const OperationEmbyFsList = "/api.emby.Emby/FsList"
const OperationEmbyGetItem = "/api.emby.Emby/GetItem"
const OperationEmbyGetItems = "/api.emby.Emby/GetItems"
const OperationEmbyGetSystemInfo = "/api.emby.Emby/GetSystemInfo"
const OperationEmbyLogin = "/api.emby.Emby/Login"
const OperationEmbyLogout = "/api.emby.Emby/Logout"
const OperationEmbyMe = "/api.emby.Emby/Me"
const OperationEmbyPlaybackInfo = "/api.emby.Emby/PlaybackInfo"

type EmbyHTTPServer interface {
	DeleteActiveEncodeings(context.Context, *DeleteActiveEncodeingsReq) (*Empty, error)
	FsList(context.Context, *FsListReq) (*FsListResp, error)
	GetItem(context.Context, *GetItemReq) (*Item, error)
	GetItems(context.Context, *GetItemsReq) (*GetItemsResp, error)
	GetSystemInfo(context.Context, *SystemInfoReq) (*SystemInfoResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Logout(context.Context, *LogoutReq) (*Empty, error)
	Me(context.Context, *MeReq) (*MeResp, error)
	PlaybackInfo(context.Context, *PlaybackInfoReq) (*PlaybackInfoResp, error)
}

func RegisterEmbyHTTPServer(s *http.Server, srv EmbyHTTPServer) {
	r := s.Route("/")
	r.POST("/emby/Login", _Emby_Login1_HTTP_Handler(srv))
	r.POST("/emby/Users/Me", _Emby_Me1_HTTP_Handler(srv))
	r.POST("/emby/Items", _Emby_GetItems0_HTTP_Handler(srv))
	r.POST("/emby/Items/{itemId}", _Emby_GetItem0_HTTP_Handler(srv))
	r.POST("/emby/System/Info", _Emby_GetSystemInfo0_HTTP_Handler(srv))
	r.POST("/emby/FileSystem/Paths", _Emby_FsList1_HTTP_Handler(srv))
	r.POST("/emby/Sessions/Logout", _Emby_Logout0_HTTP_Handler(srv))
	r.POST("/emby/Playback/Info", _Emby_PlaybackInfo0_HTTP_Handler(srv))
	r.POST("/emby/DeleteActiveEncodeing", _Emby_DeleteActiveEncodeings0_HTTP_Handler(srv))
}

func _Emby_Login1_HTTP_Handler(srv EmbyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmbyLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResp)
		return ctx.Result(200, reply)
	}
}

func _Emby_Me1_HTTP_Handler(srv EmbyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmbyMe)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Me(ctx, req.(*MeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MeResp)
		return ctx.Result(200, reply)
	}
}

func _Emby_GetItems0_HTTP_Handler(srv EmbyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetItemsReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmbyGetItems)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetItems(ctx, req.(*GetItemsReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetItemsResp)
		return ctx.Result(200, reply)
	}
}

func _Emby_GetItem0_HTTP_Handler(srv EmbyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetItemReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmbyGetItem)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetItem(ctx, req.(*GetItemReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Item)
		return ctx.Result(200, reply)
	}
}

func _Emby_GetSystemInfo0_HTTP_Handler(srv EmbyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SystemInfoReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmbyGetSystemInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSystemInfo(ctx, req.(*SystemInfoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SystemInfoResp)
		return ctx.Result(200, reply)
	}
}

func _Emby_FsList1_HTTP_Handler(srv EmbyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FsListReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmbyFsList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FsList(ctx, req.(*FsListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FsListResp)
		return ctx.Result(200, reply)
	}
}

func _Emby_Logout0_HTTP_Handler(srv EmbyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmbyLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _Emby_PlaybackInfo0_HTTP_Handler(srv EmbyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlaybackInfoReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmbyPlaybackInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PlaybackInfo(ctx, req.(*PlaybackInfoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlaybackInfoResp)
		return ctx.Result(200, reply)
	}
}

func _Emby_DeleteActiveEncodeings0_HTTP_Handler(srv EmbyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteActiveEncodeingsReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmbyDeleteActiveEncodeings)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteActiveEncodeings(ctx, req.(*DeleteActiveEncodeingsReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

type EmbyHTTPClient interface {
	DeleteActiveEncodeings(ctx context.Context, req *DeleteActiveEncodeingsReq, opts ...http.CallOption) (rsp *Empty, err error)
	FsList(ctx context.Context, req *FsListReq, opts ...http.CallOption) (rsp *FsListResp, err error)
	GetItem(ctx context.Context, req *GetItemReq, opts ...http.CallOption) (rsp *Item, err error)
	GetItems(ctx context.Context, req *GetItemsReq, opts ...http.CallOption) (rsp *GetItemsResp, err error)
	GetSystemInfo(ctx context.Context, req *SystemInfoReq, opts ...http.CallOption) (rsp *SystemInfoResp, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginResp, err error)
	Logout(ctx context.Context, req *LogoutReq, opts ...http.CallOption) (rsp *Empty, err error)
	Me(ctx context.Context, req *MeReq, opts ...http.CallOption) (rsp *MeResp, err error)
	PlaybackInfo(ctx context.Context, req *PlaybackInfoReq, opts ...http.CallOption) (rsp *PlaybackInfoResp, err error)
}

type EmbyHTTPClientImpl struct {
	cc *http.Client
}

func NewEmbyHTTPClient(client *http.Client) EmbyHTTPClient {
	return &EmbyHTTPClientImpl{client}
}

func (c *EmbyHTTPClientImpl) DeleteActiveEncodeings(ctx context.Context, in *DeleteActiveEncodeingsReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/emby/DeleteActiveEncodeing"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmbyDeleteActiveEncodeings))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EmbyHTTPClientImpl) FsList(ctx context.Context, in *FsListReq, opts ...http.CallOption) (*FsListResp, error) {
	var out FsListResp
	pattern := "/emby/FileSystem/Paths"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmbyFsList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EmbyHTTPClientImpl) GetItem(ctx context.Context, in *GetItemReq, opts ...http.CallOption) (*Item, error) {
	var out Item
	pattern := "/emby/Items/{itemId}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmbyGetItem))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EmbyHTTPClientImpl) GetItems(ctx context.Context, in *GetItemsReq, opts ...http.CallOption) (*GetItemsResp, error) {
	var out GetItemsResp
	pattern := "/emby/Items"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmbyGetItems))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EmbyHTTPClientImpl) GetSystemInfo(ctx context.Context, in *SystemInfoReq, opts ...http.CallOption) (*SystemInfoResp, error) {
	var out SystemInfoResp
	pattern := "/emby/System/Info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmbyGetSystemInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EmbyHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginResp, error) {
	var out LoginResp
	pattern := "/emby/Login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmbyLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EmbyHTTPClientImpl) Logout(ctx context.Context, in *LogoutReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/emby/Sessions/Logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmbyLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EmbyHTTPClientImpl) Me(ctx context.Context, in *MeReq, opts ...http.CallOption) (*MeResp, error) {
	var out MeResp
	pattern := "/emby/Users/Me"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmbyMe))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EmbyHTTPClientImpl) PlaybackInfo(ctx context.Context, in *PlaybackInfoReq, opts ...http.CallOption) (*PlaybackInfoResp, error) {
	var out PlaybackInfoResp
	pattern := "/emby/Playback/Info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmbyPlaybackInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
