// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: conf.proto

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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Consul        *Registry_Consul       `protobuf:"bytes,1,opt,name=consul,proto3" json:"consul,omitempty"`
	Etcd          *Registry_Etcd         `protobuf:"bytes,2,opt,name=etcd,proto3" json:"etcd,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Registry) Reset() {
	*x = Registry{}
	mi := &file_conf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[0]
	if x != nil {
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
	return file_conf_proto_rawDescGZIP(), []int{0}
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

type WebServer struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Addr           string                 `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty" env:"WEB_ADDR"` // @gotags: env:"WEB_ADDR"
	Tls            *WebServer_TLS         `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty"`
	CustomEndpoint string                 `protobuf:"bytes,3,opt,name=custom_endpoint,json=customEndpoint,proto3" json:"custom_endpoint,omitempty" env:"WEB_CUSTOM_ENDPOINT"` // @gotags: env:"WEB_CUSTOM_ENDPOINT"
	ServiceName    string                 `protobuf:"bytes,6,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty" env:"SERVICE_NAME"`          // @gotags: env:"SERVICE_NAME"
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *WebServer) Reset() {
	*x = WebServer{}
	mi := &file_conf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebServer) ProtoMessage() {}

func (x *WebServer) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebServer.ProtoReflect.Descriptor instead.
func (*WebServer) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{1}
}

func (x *WebServer) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *WebServer) GetTls() *WebServer_TLS {
	if x != nil {
		return x.Tls
	}
	return nil
}

func (x *WebServer) GetCustomEndpoint() string {
	if x != nil {
		return x.CustomEndpoint
	}
	return ""
}

func (x *WebServer) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type GrpcServer struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Addr           string                 `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty" env:"GRPC_ADDR"` // @gotags: env:"GRPC_ADDR"
	Tls            *GrpcServer_TLS        `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty"`
	CustomEndpoint string                 `protobuf:"bytes,3,opt,name=custom_endpoint,json=customEndpoint,proto3" json:"custom_endpoint,omitempty" env:"GRPC_CUSTOM_ENDPOINT"` // @gotags: env:"GRPC_CUSTOM_ENDPOINT"
	Timeout        *durationpb.Duration   `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty" env:"SERVER_TIMEOUT"`                                     // @gotags: env:"SERVER_TIMEOUT"
	JwtSecret      string                 `protobuf:"bytes,5,opt,name=jwt_secret,json=jwtSecret,proto3" json:"jwt_secret,omitempty" env:"SERVER_JWT_SECRET"`                // @gotags: env:"SERVER_JWT_SECRET"
	ServiceName    string                 `protobuf:"bytes,6,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty" env:"SERVICE_NAME"`          // @gotags: env:"SERVICE_NAME"
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GrpcServer) Reset() {
	*x = GrpcServer{}
	mi := &file_conf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrpcServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcServer) ProtoMessage() {}

func (x *GrpcServer) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcServer.ProtoReflect.Descriptor instead.
func (*GrpcServer) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2}
}

func (x *GrpcServer) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *GrpcServer) GetTls() *GrpcServer_TLS {
	if x != nil {
		return x.Tls
	}
	return nil
}

func (x *GrpcServer) GetCustomEndpoint() string {
	if x != nil {
		return x.CustomEndpoint
	}
	return ""
}

func (x *GrpcServer) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *GrpcServer) GetJwtSecret() string {
	if x != nil {
		return x.JwtSecret
	}
	return ""
}

func (x *GrpcServer) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type BilibiliConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BilibiliConfig) Reset() {
	*x = BilibiliConfig{}
	mi := &file_conf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BilibiliConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BilibiliConfig) ProtoMessage() {}

func (x *BilibiliConfig) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[3]
	if x != nil {
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
	return file_conf_proto_rawDescGZIP(), []int{3}
}

type AlistConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlistConfig) Reset() {
	*x = AlistConfig{}
	mi := &file_conf_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlistConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlistConfig) ProtoMessage() {}

func (x *AlistConfig) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlistConfig.ProtoReflect.Descriptor instead.
func (*AlistConfig) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{4}
}

type EmbyConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmbyConfig) Reset() {
	*x = EmbyConfig{}
	mi := &file_conf_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmbyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbyConfig) ProtoMessage() {}

func (x *EmbyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbyConfig.ProtoReflect.Descriptor instead.
func (*EmbyConfig) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{5}
}

type WebdavConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WebdavConfig) Reset() {
	*x = WebdavConfig{}
	mi := &file_conf_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebdavConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebdavConfig) ProtoMessage() {}

func (x *WebdavConfig) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebdavConfig.ProtoReflect.Descriptor instead.
func (*WebdavConfig) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{6}
}

type AllServer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Server        *GrpcServer            `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Registry      *Registry              `protobuf:"bytes,2,opt,name=registry,proto3" json:"registry,omitempty"`
	Bilibili      *BilibiliConfig        `protobuf:"bytes,3,opt,name=bilibili,proto3" json:"bilibili,omitempty"`
	Alist         *AlistConfig           `protobuf:"bytes,4,opt,name=alist,proto3" json:"alist,omitempty"`
	Emby          *EmbyConfig            `protobuf:"bytes,5,opt,name=emby,proto3" json:"emby,omitempty"`
	Webdav        *WebdavConfig          `protobuf:"bytes,6,opt,name=webdav,proto3" json:"webdav,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllServer) Reset() {
	*x = AllServer{}
	mi := &file_conf_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllServer) ProtoMessage() {}

func (x *AllServer) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllServer.ProtoReflect.Descriptor instead.
func (*AllServer) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{7}
}

func (x *AllServer) GetServer() *GrpcServer {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *AllServer) GetRegistry() *Registry {
	if x != nil {
		return x.Registry
	}
	return nil
}

func (x *AllServer) GetBilibili() *BilibiliConfig {
	if x != nil {
		return x.Bilibili
	}
	return nil
}

func (x *AllServer) GetAlist() *AlistConfig {
	if x != nil {
		return x.Alist
	}
	return nil
}

func (x *AllServer) GetEmby() *EmbyConfig {
	if x != nil {
		return x.Emby
	}
	return nil
}

func (x *AllServer) GetWebdav() *WebdavConfig {
	if x != nil {
		return x.Webdav
	}
	return nil
}

type Registry_Consul struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Scheme        string                 `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty" env:"CONSUL_SCHEME"`   // @gotags: env:"CONSUL_SCHEME"
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty" env:"CONSUL_ADDR"`       // @gotags: env:"CONSUL_ADDR"
	Prefix        string                 `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty" env:"CONSUL_PREFIX"`   // @gotags: env:"CONSUL_PREFIX"
	Timeout       *durationpb.Duration   `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty" env:"CONSUL_TIMEOUT"` // @gotags: env:"CONSUL_TIMEOUT"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Registry_Consul) Reset() {
	*x = Registry_Consul{}
	mi := &file_conf_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Registry_Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Consul) ProtoMessage() {}

func (x *Registry_Consul) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[8]
	if x != nil {
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
	return file_conf_proto_rawDescGZIP(), []int{0, 0}
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty" env:"ETCD_ENDPOINT"` // @gotags: env:"ETCD_ENDPOINT"
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" env:"ETCD_USERNAME"` // @gotags: env:"ETCD_USERNAME"
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" env:"ETCD_PASSWORD"` // @gotags: env:"ETCD_PASSWORD"
	Timeout       *durationpb.Duration   `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty" env:"ETCD_TIMEOUT"`   // @gotags: env:"ETCD_TIMEOUT"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Registry_Etcd) Reset() {
	*x = Registry_Etcd{}
	mi := &file_conf_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Registry_Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_Etcd) ProtoMessage() {}

func (x *Registry_Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[9]
	if x != nil {
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
	return file_conf_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Registry_Etcd) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
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

type WebServer_TLS struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CertFile      string                 `protobuf:"bytes,1,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty" env:"WEB_CERT_FILE"` // @gotags: env:"WEB_CERT_FILE"
	KeyFile       string                 `protobuf:"bytes,2,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty" env:"WEB_KEY_FILE"`    // @gotags: env:"WEB_KEY_FILE"
	CaFile        string                 `protobuf:"bytes,3,opt,name=ca_file,json=caFile,proto3" json:"ca_file,omitempty" env:"WEB_CA_FILE"`       // @gotags: env:"WEB_CA_FILE"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WebServer_TLS) Reset() {
	*x = WebServer_TLS{}
	mi := &file_conf_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebServer_TLS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebServer_TLS) ProtoMessage() {}

func (x *WebServer_TLS) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebServer_TLS.ProtoReflect.Descriptor instead.
func (*WebServer_TLS) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *WebServer_TLS) GetCertFile() string {
	if x != nil {
		return x.CertFile
	}
	return ""
}

func (x *WebServer_TLS) GetKeyFile() string {
	if x != nil {
		return x.KeyFile
	}
	return ""
}

func (x *WebServer_TLS) GetCaFile() string {
	if x != nil {
		return x.CaFile
	}
	return ""
}

type GrpcServer_TLS struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CertFile      string                 `protobuf:"bytes,1,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty" env:"GRPC_CERT_FILE"` // @gotags: env:"GRPC_CERT_FILE"
	KeyFile       string                 `protobuf:"bytes,2,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty" env:"GRPC_KEY_FILE"`    // @gotags: env:"GRPC_KEY_FILE"
	CaFile        string                 `protobuf:"bytes,3,opt,name=ca_file,json=caFile,proto3" json:"ca_file,omitempty" env:"GRPC_CA_FILE"`       // @gotags: env:"GRPC_CA_FILE"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GrpcServer_TLS) Reset() {
	*x = GrpcServer_TLS{}
	mi := &file_conf_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrpcServer_TLS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcServer_TLS) ProtoMessage() {}

func (x *GrpcServer_TLS) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcServer_TLS.ProtoReflect.Descriptor instead.
func (*GrpcServer_TLS) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GrpcServer_TLS) GetCertFile() string {
	if x != nil {
		return x.CertFile
	}
	return ""
}

func (x *GrpcServer_TLS) GetKeyFile() string {
	if x != nil {
		return x.KeyFile
	}
	return ""
}

func (x *GrpcServer_TLS) GetCaFile() string {
	if x != nil {
		return x.CaFile
	}
	return ""
}

var File_conf_proto protoreflect.FileDescriptor

var file_conf_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x2d, 0x0a, 0x04, 0x65, 0x74,
	0x63, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x45,
	0x74, 0x63, 0x64, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x1a, 0x81, 0x01, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x8f, 0x01,
	0x0a, 0x04, 0x45, 0x74, 0x63, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22,
	0xf0, 0x01, 0x0a, 0x09, 0x57, 0x65, 0x62, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x12, 0x2b, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x65, 0x62, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x4c, 0x53, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x56, 0x0a, 0x03, 0x54, 0x4c,
	0x53, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x46, 0x69,
	0x6c, 0x65, 0x22, 0xc6, 0x02, 0x0a, 0x0a, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x2c, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x4c, 0x53, 0x52, 0x03,
	0x74, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x77, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x77, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x56, 0x0a, 0x03, 0x54, 0x4c, 0x53, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65,
	0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x42,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x0d, 0x0a,
	0x0b, 0x41, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x0c, 0x0a, 0x0a,
	0x45, 0x6d, 0x62, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x0e, 0x0a, 0x0c, 0x57, 0x65,
	0x62, 0x64, 0x61, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xb2, 0x02, 0x0a, 0x09, 0x41,
	0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x36, 0x0a, 0x08, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x12, 0x2d, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x04, 0x65, 0x6d, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x62,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x65, 0x6d, 0x62, 0x79, 0x12, 0x30, 0x0a,
	0x06, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x65, 0x62, 0x64, 0x61,
	0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x42,
	0x1c, 0x5a, 0x1a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_proto_rawDescOnce sync.Once
	file_conf_proto_rawDescData = file_conf_proto_rawDesc
)

func file_conf_proto_rawDescGZIP() []byte {
	file_conf_proto_rawDescOnce.Do(func() {
		file_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_proto_rawDescData)
	})
	return file_conf_proto_rawDescData
}

var file_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_conf_proto_goTypes = []any{
	(*Registry)(nil),            // 0: kratos.api.Registry
	(*WebServer)(nil),           // 1: kratos.api.WebServer
	(*GrpcServer)(nil),          // 2: kratos.api.GrpcServer
	(*BilibiliConfig)(nil),      // 3: kratos.api.BilibiliConfig
	(*AlistConfig)(nil),         // 4: kratos.api.AlistConfig
	(*EmbyConfig)(nil),          // 5: kratos.api.EmbyConfig
	(*WebdavConfig)(nil),        // 6: kratos.api.WebdavConfig
	(*AllServer)(nil),           // 7: kratos.api.AllServer
	(*Registry_Consul)(nil),     // 8: kratos.api.Registry.Consul
	(*Registry_Etcd)(nil),       // 9: kratos.api.Registry.Etcd
	(*WebServer_TLS)(nil),       // 10: kratos.api.WebServer.TLS
	(*GrpcServer_TLS)(nil),      // 11: kratos.api.GrpcServer.TLS
	(*durationpb.Duration)(nil), // 12: google.protobuf.Duration
}
var file_conf_proto_depIdxs = []int32{
	8,  // 0: kratos.api.Registry.consul:type_name -> kratos.api.Registry.Consul
	9,  // 1: kratos.api.Registry.etcd:type_name -> kratos.api.Registry.Etcd
	10, // 2: kratos.api.WebServer.tls:type_name -> kratos.api.WebServer.TLS
	11, // 3: kratos.api.GrpcServer.tls:type_name -> kratos.api.GrpcServer.TLS
	12, // 4: kratos.api.GrpcServer.timeout:type_name -> google.protobuf.Duration
	2,  // 5: kratos.api.AllServer.server:type_name -> kratos.api.GrpcServer
	0,  // 6: kratos.api.AllServer.registry:type_name -> kratos.api.Registry
	3,  // 7: kratos.api.AllServer.bilibili:type_name -> kratos.api.BilibiliConfig
	4,  // 8: kratos.api.AllServer.alist:type_name -> kratos.api.AlistConfig
	5,  // 9: kratos.api.AllServer.emby:type_name -> kratos.api.EmbyConfig
	6,  // 10: kratos.api.AllServer.webdav:type_name -> kratos.api.WebdavConfig
	12, // 11: kratos.api.Registry.Consul.timeout:type_name -> google.protobuf.Duration
	12, // 12: kratos.api.Registry.Etcd.timeout:type_name -> google.protobuf.Duration
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_conf_proto_init() }
func file_conf_proto_init() {
	if File_conf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_proto_goTypes,
		DependencyIndexes: file_conf_proto_depIdxs,
		MessageInfos:      file_conf_proto_msgTypes,
	}.Build()
	File_conf_proto = out.File
	file_conf_proto_rawDesc = nil
	file_conf_proto_goTypes = nil
	file_conf_proto_depIdxs = nil
}
