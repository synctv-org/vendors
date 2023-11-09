package registry

import (
	"github.com/go-kratos/kratos/v2/registry"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/hashicorp/consul/api"
	"github.com/synctv-org/vendors/internal/conf"
)

func newConsul(c *conf.Registry_Consul) registry.Registrar {
	if c == nil {
		return nil
	}
	config := api.DefaultConfig()
	if c.Addr == "" {
		panic("consul address is empty")
	}
	config.Address = c.Addr
	if c.Scheme != "" {
		config.Scheme = c.Scheme
	}
	if c.Timeout != nil && c.Timeout.AsDuration() != 0 {
		config.WaitTime = c.Timeout.AsDuration()
	}
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	return consul.New(client)
}
