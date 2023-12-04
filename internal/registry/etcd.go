package registry

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/synctv-org/vendors/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func newEtcd(c *conf.Registry_Etcd) registry.Registrar {
	if c == nil {
		return nil
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Endpoints,
		Username:    c.Username,
		Password:    c.Password,
		DialTimeout: c.Timeout.AsDuration(),
	})
	if err != nil {
		panic(err)
	}
	return etcd.New(client)
}
