package utils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	ggrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	ghttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/synctv-org/vendors/cmd/flags"
	"github.com/synctv-org/vendors/conf"
	"github.com/zijiren233/go-colorable"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

const (
	UA = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.69`
)

var (
	noRedirectHttpClient = NoRedirectClient(&http.Client{})
)

func NoRedirectHttpClient() *http.Client {
	return noRedirectHttpClient
}

func NoRedirectClient(client *http.Client) *http.Client {
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return client
}

func MapToHttpCookies(m map[string]string) []*http.Cookie {
	var cookies = make([]*http.Cookie, 0, len(m))
	for k, v := range m {
		cookies = append(cookies, &http.Cookie{
			Name:  k,
			Value: v,
		})
	}
	return cookies
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
	gs          *ggrpc.Server
	hs          *ghttp.Server
	serviceName string
	once        sync.Once
}

func (s *GrpcGatewayServer) Start(ctx context.Context) error {
	s.once.Do(func() {
		s.hs.Handler = grpcHandlerFunc(s.gs, s.hs.Handler)
	})
	return s.hs.Start(ctx)
}

func (s *GrpcGatewayServer) Stop(ctx context.Context) error {
	return s.hs.Stop(ctx)
}

func (s *GrpcGatewayServer) Endpoint() (*url.URL, error) {
	return s.gs.Endpoint()
}

func (s *GrpcGatewayServer) GrpcRegistrar() grpc.ServiceRegistrar {
	return s.gs
}

func (s *GrpcGatewayServer) HttpRegistrar() *ghttp.Server {
	return s.hs
}

func (s *GrpcGatewayServer) Endpoints() ([]*url.URL, error) {
	ge, err := s.gs.Endpoint()
	if err != nil {
		return nil, err
	}
	he, err := s.hs.Endpoint()
	if err != nil {
		return nil, err
	}
	return []*url.URL{ge, he}, nil
}

func (s *GrpcGatewayServer) ServiceName() string {
	return s.serviceName
}

func NewGrpcGatewayServer(config *conf.GrpcServer) *GrpcGatewayServer {
	middlewares := []middleware.Middleware{recovery.Recovery()}
	if config.JwtSecret != "" {
		jwtSecret := []byte(config.JwtSecret)
		middlewares = append(middlewares, jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
			return jwtSecret, nil
		}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256)))
	}

	l, err := net.Listen("tcp", config.Addr)
	if err != nil {
		panic(err)
	}

	var hopts = []ghttp.ServerOption{
		ghttp.Middleware(middlewares...),
		ghttp.Listener(l),
		ghttp.Address(config.Addr),
	}

	var gopts = []ggrpc.ServerOption{
		ggrpc.Middleware(middlewares...),
		ggrpc.Listener(l),
		ggrpc.Address(config.Addr),
	}

	if config.Timeout != nil {
		hopts = append(hopts, ghttp.Timeout(config.Timeout.AsDuration()))
		gopts = append(gopts, ggrpc.Timeout(config.Timeout.AsDuration()))
	}

	var enableTls bool
	if config.Tls != nil && config.Tls.CertFile != "" && config.Tls.KeyFile != "" {
		enableTls = true
		var rootCAs *x509.CertPool
		rootCAs, err := x509.SystemCertPool()
		if err != nil {
			fmt.Println("Failed to load system root CA:", err)
			panic(err)
		}
		if config.Tls.CaFile != "" {
			b, err := os.ReadFile(config.Tls.CaFile)
			if err != nil {
				panic(err)
			}
			rootCAs.AppendCertsFromPEM(b)
		}

		cert, err := tls.LoadX509KeyPair(config.Tls.CertFile, config.Tls.KeyFile)
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

	if config.CustomEndpoint != "" {
		u, err := url.Parse(config.CustomEndpoint)
		if err != nil {
			panic(err)
		}
		var (
			hu = *u
			gu = *u
		)
		if u.Scheme == "grpcs" || u.Scheme == "https" {
			hu.Scheme = "https"
			gu.Scheme = "grpcs"
		} else if u.Scheme == "grpc" || u.Scheme == "http" {
			hu.Scheme = "http"
			gu.Scheme = "grpc"
		} else if u.Scheme == "" {
			if enableTls {
				hu.Scheme = "https"
				gu.Scheme = "grpcs"
			} else {
				hu.Scheme = "http"
				gu.Scheme = "grpc"
			}
		} else {
			panic("invalid custom endpoint scheme: " + u.Scheme)
		}
		hopts = append(hopts, ghttp.Endpoint(&hu))
		gopts = append(gopts, ggrpc.Endpoint(&gu))
	}

	hs := ghttp.NewServer(hopts...)
	gs := ggrpc.NewServer(gopts...)
	return &GrpcGatewayServer{
		gs: gs,
		hs: hs,
	}
}

func GetEnvFiles(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasPrefix(info.Name(), ".env") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

var (
	needColor     bool
	needColorOnce sync.Once
)

func ForceColor() bool {
	needColorOnce.Do(func() {
		if flags.DisableLogColor {
			needColor = false
			return
		}
		needColor = colorable.IsTerminal(os.Stdout.Fd())
	})
	return needColor
}
