package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/synctv-org/vendors/cmd/flags"
	"github.com/synctv-org/vendors/cmd/server"
)

var RootCmd = &cobra.Command{
	Use:   "vendors",
	Short: "synctv vendors",
	Long:  `synctv https://github.com/synctv-org/vendors`,
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
