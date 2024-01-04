package server

import (
	bilibiliApi "github.com/synctv-org/vendors/api/bilibili"
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/service/bilibili"
	"github.com/synctv-org/vendors/utils"
)

func NewGRPCServer(
	config *conf.GrpcServer,
	bilibili *bilibili.BilibiliService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	bilibiliApi.RegisterBilibiliServer(ggs.GrpcRegistrar(), bilibili)
	bilibiliApi.RegisterBilibiliHTTPServer(ggs.HttpRegistrar(), bilibili)
	return ggs
}
