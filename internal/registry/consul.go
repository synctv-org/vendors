package registry

import (
	"github.com/hashicorp/consul/api"
	"github.com/synctv-org/vendors/conf"
)

func newConsul(c *conf.Registry_Consul) *api.Client {
	if c == nil || c.Addr == "" {
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
	return client
}
