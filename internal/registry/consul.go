package registry

import (
	"github.com/hashicorp/consul/api"
	"github.com/synctv-org/vendors/conf"
)

func newConsul(c *conf.Registry_Consul) *api.Client {
	if c == nil || c.GetAddr() == "" {
		return nil
	}
	config := api.DefaultConfig()
	if c.GetAddr() == "" {
		panic("consul address is empty")
	}
	config.Address = c.GetAddr()
	if c.GetScheme() != "" {
		config.Scheme = c.GetScheme()
	}
	if c.GetTimeout() != nil && c.GetTimeout().AsDuration() != 0 {
		config.WaitTime = c.GetTimeout().AsDuration()
	}
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	return client
}
