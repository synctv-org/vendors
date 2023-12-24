package conf

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

func DefaultServer() *Server {
	return &Server{
		Grpc: &GRPC{
			Addr: ":9000",
		},
		Web: &Web{
			Addr: ":8000",
		},
		Timeout: durationpb.New(time.Second * 15),
	}
}
