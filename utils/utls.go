package utils

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	utls "github.com/refraction-networking/utls"
	"golang.org/x/net/http2"
)

var _ io.ReadCloser = (*utlsHttpBody)(nil)

type utlsHttpBody struct {
	*utls.UConn
	rawBody io.ReadCloser
}

func (u *utlsHttpBody) Read(p []byte) (n int, err error) {
	return u.rawBody.Read(p)
}

func (u *utlsHttpBody) Close() error {
	defer u.UConn.Close()
	return u.rawBody.Close()
}

type UtlsHttpOption func(*UtlsHttpRoundTripper)

func WithHttpClient(client *http.Client) UtlsHttpOption {
	return func(options *UtlsHttpRoundTripper) {
		if client != nil {
			options.Timeout = client.Timeout
			options.Base = client.Transport
		}
	}
}

func WithContext(ctx context.Context) UtlsHttpOption {
	return func(options *UtlsHttpRoundTripper) {
		options.Ctx = ctx
	}
}

func WithBaseTransport(base http.RoundTripper) UtlsHttpOption {
	return func(options *UtlsHttpRoundTripper) {
		options.Base = base
	}
}

var DefaultUtlsHttpRoundTripper = NewUtlsHttpRoundTripper()

func NewUtlsHttpRoundTripper(opts ...UtlsHttpOption) *UtlsHttpRoundTripper {
	rt := &UtlsHttpRoundTripper{}
	for _, opt := range opts {
		opt(rt)
	}
	return rt
}

func UtlsDo(req *http.Request) (*http.Response, error) {
	return DefaultUtlsHttpRoundTripper.RoundTrip(req)
}

func UtlsDoWithOptions(req *http.Request, opts ...UtlsHttpOption) (*http.Response, error) {
	return NewUtlsHttpRoundTripper(opts...).RoundTrip(req)
}

type UtlsHttpRoundTripper struct {
	Base    http.RoundTripper
	Ctx     context.Context
	Timeout time.Duration
}

func (u *UtlsHttpRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if u.Ctx == nil {
		u.Ctx = context.Background()
	}

	config := utls.Config{ServerName: req.URL.Hostname()}
	port := req.URL.Port()
	if port == "" {
		switch req.URL.Scheme {
		case "https":
			port = "443"
		case "http":
			port = "80"
		}
	}
	dialConn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", req.URL.Hostname(), port), u.Timeout)
	if err != nil {
		return nil, err
	}
	uTlsConn := utls.UClient(dialConn, &config, utls.HelloChrome_Auto)
	err = uTlsConn.HandshakeContext(u.Ctx)
	if err != nil {
		return nil, err
	}
	resp, err := httpGetOverConn(u.Base, req, uTlsConn, uTlsConn.ConnectionState().NegotiatedProtocol)
	if err != nil {
		_ = uTlsConn.Close()
		return nil, err
	}
	resp.Body = &utlsHttpBody{uTlsConn, resp.Body}
	return resp, nil
}

func getH2RoundTripper(rt http.RoundTripper, conn net.Conn) (http.RoundTripper, error) {
	if rt != nil {
		tr, ok := rt.(*http.Transport)
		if !ok {
			tr = tr.Clone()
			tr.MaxResponseHeaderBytes = 262144
			tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
				return conn, nil
			}
			h2transport, err := http2.ConfigureTransports(tr)
			if err != nil {
				return nil, err
			}
			c, err := h2transport.NewClientConn(conn)
			if err != nil {
				return nil, err
			}
			rt = c
		} else if h2tr, ok := rt.(*http2.Transport); ok {
			h2tr.MaxHeaderListSize = 262144
			c, err := h2tr.NewClientConn(conn)
			if err != nil {
				return nil, err
			}
			rt = c
		} else {
			return nil, fmt.Errorf("unsupported RoundTripper: %T", rt)
		}
	} else {
		tr := http2.Transport{MaxHeaderListSize: 262144}
		c, err := tr.NewClientConn(conn)
		if err != nil {
			return nil, err
		}
		rt = c
	}
	return rt, nil
}

func getH1_1RoundTripper(rt http.RoundTripper, conn net.Conn) (http.RoundTripper, error) {
	if rt != nil {
		tr, ok := rt.(*http.Transport)
		if !ok {
			return nil, fmt.Errorf("unsupported RoundTripper: %T", rt)
		}
		tr = tr.Clone()
		tr.MaxResponseHeaderBytes = 262144
		tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return conn, nil
		}
		rt = tr
	} else {
		tr := (http.DefaultTransport).(*http.Transport).Clone()
		tr.MaxResponseHeaderBytes = 262144
		tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return conn, nil
		}
		rt = tr
	}
	return rt, nil
}

func httpGetOverConn(rt http.RoundTripper, req *http.Request, conn net.Conn, alpn string) (*http.Response, error) {
	var err error
	switch alpn {
	case "h2":
		req.Proto = "HTTP/2.0"
		req.ProtoMajor = 2
		req.ProtoMinor = 0
		rt, err = getH2RoundTripper(rt, conn)
		if err != nil {
			return nil, err
		}
	case "http/1.1", "":
		req.Proto = "HTTP/1.1"
		req.ProtoMajor = 1
		req.ProtoMinor = 1
		rt, err = getH1_1RoundTripper(rt, conn)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported ALPN: %v", alpn)
	}
	return rt.RoundTrip(req)
}
