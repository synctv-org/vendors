package server

import (
	embyApi "github.com/synctv-org/vendors/api/emby"
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/service/emby"
	"github.com/synctv-org/vendors/utils"
)

func NewGRPCServer(
	config *conf.GrpcServer,
	emby *emby.EmbyService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	embyApi.RegisterEmbyServer(ggs.GrpcRegistrar(), emby)
	embyApi.RegisterEmbyHTTPServer(ggs.HttpRegistrar(), emby)
	return ggs
}
