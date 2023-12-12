package utils

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	ggrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	ghttp "github.com/go-kratos/kratos/v2/transport/http"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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
	gs *ggrpc.Server
	hs *ghttp.Server
}

func NewGrpcGatewayServer(gs *ggrpc.Server, hs *ghttp.Server) *GrpcGatewayServer {
	return &GrpcGatewayServer{
		gs: gs,
		hs: hs,
	}
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
