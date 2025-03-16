package registry

import (
	"strings"

	"github.com/synctv-org/vendors/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func newEtcd(c *conf.Registry_Etcd) *clientv3.Client {
	if c == nil || c.GetEndpoint() == "" {
		return nil
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(c.GetEndpoint(), ","),
		Username:    c.GetUsername(),
		Password:    c.GetPassword(),
		DialTimeout: c.GetTimeout().AsDuration(),
	})
	if err != nil {
		panic(err)
	}
	return client
}
