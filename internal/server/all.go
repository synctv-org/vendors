package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/url"
	"os"

	jwtv4 "github.com/golang-jwt/jwt/v4"
	alistApi "github.com/synctv-org/vendors/api/alist"
	bilibiliApi "github.com/synctv-org/vendors/api/bilibili"
	embyApi "github.com/synctv-org/vendors/api/emby"
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/service/alist"
	"github.com/synctv-org/vendors/service/bilibili"
	"github.com/synctv-org/vendors/service/emby"
	"github.com/synctv-org/vendors/utils"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	ggrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	ghttp "github.com/go-kratos/kratos/v2/transport/http"
)

func NewGRPCServer(
	c *conf.Server,
	bilibili *bilibili.BilibiliService,
	alist *alist.AlistService,
	emby *emby.EmbyService,
	logger log.Logger,
) *utils.GrpcGatewayServer {
	middlewares := []middleware.Middleware{recovery.Recovery()}
	if c.JwtSecret != "" {
		jwtSecret := []byte(c.JwtSecret)
		middlewares = append(middlewares, jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
			return jwtSecret, nil
		}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)))
	}
	var hopts = []ghttp.ServerOption{
		ghttp.Middleware(middlewares...),
		ghttp.Address(c.Grpc.Addr),
	}
	if c.Timeout != nil {
		hopts = append(hopts, ghttp.Timeout(c.Timeout.AsDuration()))
	}
	if c.Grpc.CustomEndpoint != "" {
		u, err := url.Parse(c.Grpc.CustomEndpoint)
		if err != nil {
			panic(err)
		}
		hopts = append(hopts, ghttp.Endpoint(u))
	}

	var gopts = []ggrpc.ServerOption{
		ggrpc.Middleware(middlewares...),
		ggrpc.Address(c.Grpc.Addr),
	}

	if c.Grpc.Tls != nil && c.Grpc.Tls.CertFile != "" && c.Grpc.Tls.KeyFile != "" {
		var rootCAs *x509.CertPool
		rootCAs, err := x509.SystemCertPool()
		if err != nil {
			fmt.Println("Failed to load system root CA:", err)
			panic(err)
		}
		if c.Grpc.Tls.CaFile != "" {
			b, err := os.ReadFile(c.Grpc.Tls.CaFile)
			if err != nil {
				panic(err)
			}
			rootCAs.AppendCertsFromPEM(b)
		}

		cert, err := tls.LoadX509KeyPair(c.Grpc.Tls.CertFile, c.Grpc.Tls.KeyFile)
		if err != nil {
			panic(err)
		}
		hopts = append(hopts, ghttp.TLSConfig(&tls.Config{
			RootCAs:      rootCAs,
			Certificates: []tls.Certificate{cert},
		}))
		gopts = append(gopts, ggrpc.TLSConfig(&tls.Config{
			RootCAs:      rootCAs,
			Certificates: []tls.Certificate{cert},
		}))
	}

	hs := ghttp.NewServer(hopts...)

	if c.Timeout != nil {
		gopts = append(gopts, ggrpc.Timeout(c.Timeout.AsDuration()))
	}
	if c.Grpc.CustomEndpoint != "" {
		u, err := url.Parse(c.Grpc.CustomEndpoint)
		if err != nil {
			panic(err)
		}
		gopts = append(gopts, ggrpc.Endpoint(u))
	} else {
		// fix streamServerInterceptor panic(endpoint is nil)
		var (
			u   *url.URL
			err error
		)
		if c.Grpc.Tls != nil {
			u, err = url.Parse("grpcs://" + c.Grpc.Addr)
		} else {
			u, err = url.Parse("grpc://" + c.Grpc.Addr)
		}
		if err != nil {
			panic(err)
		}
		gopts = append(gopts, ggrpc.Endpoint(u))
	}

	gs := ggrpc.NewServer(gopts...)
	alistApi.RegisterAlistServer(gs, alist)
	alistApi.RegisterAlistHTTPServer(hs, alist)
	bilibiliApi.RegisterBilibiliServer(gs, bilibili)
	bilibiliApi.RegisterBilibiliHTTPServer(hs, bilibili)
	embyApi.RegisterEmbyServer(gs, emby)
	embyApi.RegisterEmbyHTTPServer(hs, emby)
	return utils.NewGrpcGatewayServer(gs, hs)
}
