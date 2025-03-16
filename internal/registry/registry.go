package registry

import (
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"github.com/synctv-org/vendors/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(NewRegistry)

func NewRegistry(reg *conf.Registry) registry.Registrar {
	if reg == nil {
		log.Infof("no registry configed")
		return nil
	}
	if reg.GetConsul() != nil && reg.GetConsul().GetAddr() != "" {
		log.Infof("use consul: %v", reg.GetConsul())
		return newConsulRegistry(newConsul(reg.GetConsul()))
	}
	if reg.GetEtcd() != nil && reg.GetEtcd().GetEndpoint() != "" {
		log.Infof("use etcd: %v", reg.GetEtcd())
		return newEtcdRegistry(newEtcd(reg.GetEtcd()))
	}
	log.Infof("no registry configed")
	return nil
}

type EtcdRegistry struct {
	client *clientv3.Client
	*etcd.Registry
}

func (e *EtcdRegistry) Client() *clientv3.Client {
	return e.client
}

func newEtcdRegistry(client *clientv3.Client) *EtcdRegistry {
	return &EtcdRegistry{
		client:   client,
		Registry: etcd.New(client),
	}
}

type ConsulRegistry struct {
	client *api.Client
	*consul.Registry
}

func (c *ConsulRegistry) Client() *api.Client {
	return c.client
}

func newConsulRegistry(client *api.Client) *ConsulRegistry {
	return &ConsulRegistry{
		client:   client,
		Registry: consul.New(client),
	}
}
