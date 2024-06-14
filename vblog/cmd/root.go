package cmd

import (
	"fmt"

	initCmd "github.com/qinchi-ops/govlog/vblog/cmd/init"
	"github.com/qinchi-ops/govlog/vblog/cmd/start"
	"github.com/qinchi-ops/govlog/vblog/conf"
	"github.com/spf13/cobra"
)

var (
	configPath string
)

var RootCmd = &cobra.Command{
	Use:   "vblog",
	Short: "vblog service",
	Run: func(cmd *cobra.Command, args []string) {

		// 1. 加载配置
		cobra.CheckErr(conf.LoadConfigFromYaml(configPath))

		// vblog version
		// v0.0.1
		if len(args) > 0 {
			if args[0] == "verson" {
				fmt.Println("v0.0.1")
			}
		} else {
			cmd.Help()
		}
	},
}

func init() {
	// --config
	RootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "etc/application.yaml", "the service config file")

	// Root--> init
	RootCmd.AddCommand(initCmd.Cmd)
	// Root --> start
	RootCmd.AddCommand(start.Cmd)
}
