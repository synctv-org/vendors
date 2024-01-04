package conf

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

func DefaultGrpcServer() *GrpcServer {
	return &GrpcServer{
		Addr:    ":9000",
		Timeout: durationpb.New(time.Second * 15),
	}
}

func DefaultRegistry() *Registry {
	return &Registry{
		Consul: &Registry_Consul{},
		Etcd:   &Registry_Etcd{},
	}
}
