package server

import (
	alistApi "github.com/synctv-org/vendors/api/alist"
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/service/alist"
	"github.com/synctv-org/vendors/utils"
)

func NewGRPCServer(
	config *conf.Server,
	alist *alist.AlistService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	alistApi.RegisterAlistServer(ggs.GrpcRegistrar(), alist)
	alistApi.RegisterAlistHTTPServer(ggs.HttpRegistrar(), alist)
	return ggs
}
