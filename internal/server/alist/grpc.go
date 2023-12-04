package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	jwtv4 "github.com/golang-jwt/jwt/v4"
	bili "github.com/synctv-org/vendors/api/alist"
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/service/alist"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	ggrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	ghttp "github.com/go-kratos/kratos/v2/transport/http"
)

func NewGRPCServer(
	c *conf.Server,
	alist *alist.AlistService,
	logger log.Logger,
) *GrpcGatewayServer {
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
	bili.RegisterAlistServer(gs, alist)
	bili.RegisterAlistHTTPServer(hs, alist)
	return &GrpcGatewayServer{
		gs: gs,
		hs: hs,
	}
}

func grpcHandlerFunc(gs *ggrpc.Server, other http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			gs.ServeHTTP(w, r)
		} else {
			other.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

type GrpcGatewayServer struct {
	gs *ggrpc.Server
	hs *ghttp.Server
}

func (s *GrpcGatewayServer) Start(ctx context.Context) error {
	s.hs.Handler = grpcHandlerFunc(s.gs, s.hs.Handler)
	return s.hs.Start(ctx)
}

func (s *GrpcGatewayServer) Stop(ctx context.Context) error {
	return s.hs.Stop(ctx)
}

func (s *GrpcGatewayServer) Endpoint() (*url.URL, error) {
	return s.hs.Endpoint()
}
