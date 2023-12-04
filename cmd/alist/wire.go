//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/synctv-org/vendors/conf"
	reg "github.com/synctv-org/vendors/internal/registry"
	server "github.com/synctv-org/vendors/internal/server/alist"
	"github.com/synctv-org/vendors/service/alist"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.Server, *conf.Registry, *conf.AlistConfig, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, alist.ProviderSet, reg.ProviderSet, newApp))
}
