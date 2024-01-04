package registry

import (
	"strings"

	"github.com/synctv-org/vendors/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func newEtcd(c *conf.Registry_Etcd) *clientv3.Client {
	if c == nil || c.Endpoint == "" {
		return nil
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(c.Endpoint, ","),
		Username:    c.Username,
		Password:    c.Password,
		DialTimeout: c.Timeout.AsDuration(),
	})
	if err != nil {
		panic(err)
	}
	return client
}
