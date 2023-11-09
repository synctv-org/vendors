package server

import (
	"net/url"

	bili "github.com/synctv-org/vendors/api/bilibili"
	"github.com/synctv-org/vendors/internal/conf"
	"github.com/synctv-org/vendors/service/bilibili"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

func NewHTTPServer(
	c *conf.Server,
	bilibili *bilibili.BilibiliService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(c.JwtSecret), nil
			}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	if c.Http.CustomEndpoint != "" {
		u, err := url.Parse(c.Http.CustomEndpoint)
		if err != nil {
			panic(err)
		}
		opts = append(opts, http.Endpoint(u))
	}
	srv := http.NewServer(opts...)
	bili.RegisterBilibiliHTTPServer(srv, bilibili)
	return srv
}
