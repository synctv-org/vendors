// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/internal/registry"
	"github.com/synctv-org/vendors/internal/server"
	"github.com/synctv-org/vendors/service/alist"
	"github.com/synctv-org/vendors/service/bilibili"
	"github.com/synctv-org/vendors/service/emby"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func wireApp(confServer *conf.GrpcServer, confRegistry *conf.Registry, alistConfig *conf.AlistConfig, bilibiliConfig *conf.BilibiliConfig, embyConfig *conf.EmbyConfig, logger log.Logger) (*kratos.App, func(), error) {
	bilibiliService := bilibili.NewBilibiliService(bilibiliConfig)
	alistService := alist.NewAlistService(alistConfig)
	embyService := emby.NewEmbyService(embyConfig)
	grpcGatewayServer := server.NewGRPCServer(confServer, bilibiliService, alistService, embyService)
	registrar := registry.NewRegistry(confRegistry)
	app := newApp(logger, grpcGatewayServer, registrar)
	return app, func() {
	}, nil
}
