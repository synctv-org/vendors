package server

import (
	alistApi "github.com/synctv-org/vendors/api/alist"
	bilibiliApi "github.com/synctv-org/vendors/api/bilibili"
	embyApi "github.com/synctv-org/vendors/api/emby"
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/service/alist"
	"github.com/synctv-org/vendors/service/bilibili"
	"github.com/synctv-org/vendors/service/emby"
	"github.com/synctv-org/vendors/utils"
)

func NewGRPCServer(
	config *conf.GrpcServer,
	bilibili *bilibili.BilibiliService,
	alist *alist.AlistService,
	emby *emby.EmbyService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	gr := ggs.GrpcRegistrar()
	hr := ggs.HttpRegistrar()
	alistApi.RegisterAlistServer(gr, alist)
	alistApi.RegisterAlistHTTPServer(hr, alist)
	bilibiliApi.RegisterBilibiliServer(gr, bilibili)
	bilibiliApi.RegisterBilibiliHTTPServer(hr, bilibili)
	embyApi.RegisterEmbyServer(gr, emby)
	embyApi.RegisterEmbyHTTPServer(hr, emby)
	return ggs
}
