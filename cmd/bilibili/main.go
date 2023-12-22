package main

import (
	"flag"
	"net/url"
	"os"

	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/utils"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&Name, "name", "", "server name")
}

func newApp(logger log.Logger, gs *utils.GrpcGatewayServer, hs *http.Server, r registry.Registrar) *kratos.App {
	s := make([]transport.Server, 0, 2)
	if gs != nil {
		s = append(s, gs)
	}
	if hs != nil {
		s = append(s, hs)
	}
	es := make([]*url.URL, 0, 2)
	if gs != nil {
		ges, err := gs.Endpoints()
		if err != nil {
			panic(err)
		}
		es = append(es, ges...)
	}
	if hs != nil {
		he, err := hs.Endpoint()
		if err != nil {
			panic(err)
		}
		es = append(es, he)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			s...,
		),
		kratos.Registrar(r),
		kratos.Endpoint(es...),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.BilibiliServer
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Bilibili.Server, bc.Bilibili.Registry, bc.Bilibili.Config, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
