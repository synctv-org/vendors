package registry

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/synctv-org/vendors/conf"
)

var ProviderSet = wire.NewSet(NewRegistry)

func NewRegistry(reg *conf.Registry) registry.Registrar {
	if reg == nil {
		return nil
	}
	if reg.Consul != nil {
		return newConsul(reg.Consul)
	}
	if reg.Etcd != nil {
		return newEtcd(reg.Etcd)
	}
	return nil
}
