// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consul *Registry_Consul `protobuf:"bytes,1,opt,name=consul,proto3" json:"consul,omitempty"`
	Etcd   *Registry_Etcd   `protobuf:"bytes,2,opt,name=etcd,proto3" json:"etcd,omitempty"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Registry) GetConsul() *Registry_Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *Registry) GetEtcd() *Registry_Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http      *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc      *Server_GRPC `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	JwtSecret string       `protobuf:"bytes,3,opt,name=jwt_secret,json=jwtSecret,proto3" json:"jwt_secret,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Server) GetJwtSecret() string {
	if x != nil {
		return x.JwtSecret
	}
	return ""
}

type BilibiliServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   *Server         `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Registry *Registry       `protobuf:"bytes,2,opt,name=registry,proto3" json:"registry,omitempty"`
	Config   *BilibiliConfig `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *BilibiliServer) Reset() {
	*x = BilibiliServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BilibiliServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BilibiliServer) ProtoMessage() {}

func (x *BilibiliServer) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BilibiliServer.ProtoReflect.Descriptor instead.
func (*BilibiliServer) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *BilibiliServer) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *BilibiliServer) GetRegistry() *Registry {
	if x != nil {
		return x.Registry
	}
	return nil
}

func (x *BilibiliServer) GetConfig() *BilibiliConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type BilibiliConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BilibiliConfig) Reset() {
	*x = BilibiliConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BilibiliConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BilibiliConfig) ProtoMessage() {}

func (x *BilibiliConfig) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BilibiliConfig.ProtoReflect.Descriptor instead.
func (*BilibiliConfig) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3}
}

type Bilibili struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bilibili *BilibiliServer `protobuf:"bytes,1,opt,name=bilibili,proto3" json:"bilibili,omitempty"`
}

func (x *Bilibili) Reset() {
	*x = Bilibili{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bilibili) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bilibili) ProtoMessage() {}

func (x *Bilibili) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bilibili.ProtoReflect.Descriptor instead.
func (*Bilibili) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Bilibili) GetBilibili() *BilibiliServer {
	if x != nil {
		return x.Bilibili
	}
	return nil
}

type Registry_Consul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme  string               `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Prefix  string               `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Registry_Consul) Reset() {
	*x = Registry_Consul{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry_Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Consul) ProtoMessage() {}

func (x *Registry_Consul) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Consul.ProtoReflect.Descriptor instead.
func (*Registry_Consul) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Registry_Consul) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Registry_Consul) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Registry_Consul) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *Registry_Consul) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Registry_Etcd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string             `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Username  string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password  string               `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Timeout   *durationpb.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Registry_Etcd) Reset() {
	*x = Registry_Etcd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry_Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Etcd) ProtoMessage() {}

func (x *Registry_Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_Etcd.ProtoReflect.Descriptor instead.
func (*Registry_Etcd) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Registry_Etcd) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *Registry_Etcd) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Registry_Etcd) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Registry_Etcd) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network        string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr           string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout        *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	CustomEndpoint string               `protobuf:"bytes,4,opt,name=custom_endpoint,json=customEndpoint,proto3" json:"custom_endpoint,omitempty"`
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_HTTP) GetCustomEndpoint() string {
	if x != nil {
		return x.CustomEndpoint
	}
	return ""
}

type Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network        string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr           string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout        *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	CustomEndpoint string               `protobuf:"bytes,4,opt,name=custom_endpoint,json=customEndpoint,proto3" json:"custom_endpoint,omitempty"`
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_GRPC) GetCustomEndpoint() string {
	if x != nil {
		return x.CustomEndpoint
	}
	return ""
}

var File_conf_conf_proto protoreflect.FileDescriptor

var file_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12,
	0x2d, 0x0a, 0x04, 0x65, 0x74, 0x63, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x1a, 0x81,
	0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x33, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x1a, 0x91, 0x01, 0x0a, 0x04, 0x45, 0x74, 0x63, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xab, 0x03, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x2b,
	0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x6a,
	0x77, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6a, 0x77, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x1a, 0x92, 0x01, 0x0a, 0x04, 0x48,
	0x54, 0x54, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a,
	0x92, 0x01, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x42, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x42, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x10, 0x0a, 0x0e, 0x42, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x42, 0x0a, 0x08, 0x42,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x12, 0x36, 0x0a, 0x08, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x08, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x42,
	0x1c, 0x5a, 0x1a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_conf_proto_rawDescOnce sync.Once
	file_conf_conf_proto_rawDescData = file_conf_conf_proto_rawDesc
)

func file_conf_conf_proto_rawDescGZIP() []byte {
	file_conf_conf_proto_rawDescOnce.Do(func() {
		file_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_conf_proto_rawDescData)
	})
	return file_conf_conf_proto_rawDescData
}

var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_conf_conf_proto_goTypes = []interface{}{
	(*Registry)(nil),            // 0: kratos.api.Registry
	(*Server)(nil),              // 1: kratos.api.Server
	(*BilibiliServer)(nil),      // 2: kratos.api.BilibiliServer
	(*BilibiliConfig)(nil),      // 3: kratos.api.BilibiliConfig
	(*Bilibili)(nil),            // 4: kratos.api.Bilibili
	(*Registry_Consul)(nil),     // 5: kratos.api.Registry.Consul
	(*Registry_Etcd)(nil),       // 6: kratos.api.Registry.Etcd
	(*Server_HTTP)(nil),         // 7: kratos.api.Server.HTTP
	(*Server_GRPC)(nil),         // 8: kratos.api.Server.GRPC
	(*durationpb.Duration)(nil), // 9: google.protobuf.Duration
}
var file_conf_conf_proto_depIdxs = []int32{
	5,  // 0: kratos.api.Registry.consul:type_name -> kratos.api.Registry.Consul
	6,  // 1: kratos.api.Registry.etcd:type_name -> kratos.api.Registry.Etcd
	7,  // 2: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
	8,  // 3: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
	1,  // 4: kratos.api.BilibiliServer.server:type_name -> kratos.api.Server
	0,  // 5: kratos.api.BilibiliServer.registry:type_name -> kratos.api.Registry
	3,  // 6: kratos.api.BilibiliServer.config:type_name -> kratos.api.BilibiliConfig
	2,  // 7: kratos.api.Bilibili.bilibili:type_name -> kratos.api.BilibiliServer
	9,  // 8: kratos.api.Registry.Consul.timeout:type_name -> google.protobuf.Duration
	9,  // 9: kratos.api.Registry.Etcd.timeout:type_name -> google.protobuf.Duration
	9,  // 10: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	9,  // 11: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_conf_conf_proto_init() }
func file_conf_conf_proto_init() {
	if File_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BilibiliServer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BilibiliConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bilibili); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry_Consul); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry_Etcd); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_conf_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_GRPC); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_conf_proto_goTypes,
		DependencyIndexes: file_conf_conf_proto_depIdxs,
		MessageInfos:      file_conf_conf_proto_msgTypes,
	}.Build()
	File_conf_conf_proto = out.File
	file_conf_conf_proto_rawDesc = nil
	file_conf_conf_proto_goTypes = nil
	file_conf_conf_proto_depIdxs = nil
}
