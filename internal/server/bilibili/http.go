package server

import (
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/service/bilibili"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(
	c *conf.Server,
	bilibili *bilibili.BilibiliService,
	logger log.Logger,
) *http.Server {
	return nil
}
