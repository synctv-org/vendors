//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package server

import (
	"github.com/synctv-org/vendors/conf"
	reg "github.com/synctv-org/vendors/internal/registry"
	"github.com/synctv-org/vendors/internal/server"
	"github.com/synctv-org/vendors/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.GrpcServer, *conf.Registry, *conf.AlistConfig, *conf.BilibiliConfig, *conf.EmbyConfig, *conf.WebdavConfig, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, reg.ProviderSet, newApp))
}
