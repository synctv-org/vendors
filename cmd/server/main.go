package server

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/synctv-org/vendors/cmd/flags"
	"github.com/synctv-org/vendors/conf"
	"github.com/synctv-org/vendors/utils"

	"github.com/caarlos0/env/v9"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"

	_ "go.uber.org/automaxprocs"
)

var (
	flagconf string

	id, _ = os.Hostname()
)

func newApp(logger log.Logger, s *utils.GrpcGatewayServer, r registry.Registrar) *kratos.App {
	es, err := s.Endpoints()
	if err != nil {
		panic(err)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name(s.ServiceName()),
		kratos.Version(flags.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			s,
		),
		kratos.Registrar(r),
		kratos.Endpoint(es...),
	)
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start synctv vendors server",
	Run:   Server,
}

func Server(cmd *cobra.Command, args []string) {
	var bc conf.AllServer = conf.AllServer{
		All: &conf.AllServer_All{
			Server: conf.DefaultServer(),
		},
	}

	if flagconf != "" {
		c := config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
		defer c.Close()

		if err := c.Load(); err != nil {
			panic(err)
		}
		if err := c.Scan(&bc); err != nil {
			panic(err)
		}
	}

	if err := env.Parse(bc.All); err != nil {
		panic(err)
	}

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", bc.All.Server.ServiceName,
		"service.version", flags.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	app, cleanup, err := wireApp(bc.All.Server, bc.All.Registry, bc.All.Alist, bc.All.Bilibili, bc.All.Emby, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func init() {
	ServerCmd.PersistentFlags().StringVarP(&flagconf, "conf", "c", "", "config path, eg: -c config.yaml")
}