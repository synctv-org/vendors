package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/synctv-org/vendors/cmd/flags"
	"github.com/synctv-org/vendors/cmd/server"
	"github.com/synctv-org/vendors/internal/bootstrap"
	"github.com/synctv-org/vendors/utils"
)

var RootCmd = &cobra.Command{
	Use:   "vendors",
	Short: "synctv vendors",
	Long:  `synctv https://github.com/synctv-org/vendors`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		s, err := utils.GetEnvFiles(".")
		if err != nil {
			logrus.Fatalf("error getting .env files: %v", err)
		}
		if len(s) != 0 {
			err = godotenv.Overload(s...)
			if err != nil {
				logrus.Fatalf("error loading .env files: %v", err)
			}
		}
		_ = bootstrap.InitLog()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(server.ServerCmd)
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&flags.Dev, "dev", "d", false, "dev mode")
}
