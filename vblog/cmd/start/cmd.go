package start

import (
	"os"

	"github.com/qinchi-ops/govlog/vblog/conf"
	"github.com/qinchi-ops/govlog/vblog/ioc"
	"github.com/spf13/cobra"

	//业务对象注册
	_ "github.com/qinchi-ops/govlog/vblog/apps"
)

var (
	testParam string
)

var Cmd = &cobra.Command{
	Use:   "start",
	Short: "start vblog api server",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		// 1. 加载配置
		configPath := os.Getenv("CONFIG_PATH")
		if configPath == "" {
			configPath = "etc/application.yaml"
		}
		if err := conf.LoadConfigFromYaml(configPath); err != nil {
			panic(err)
		}

		// 2. 初始化Ioc
		if err := ioc.Controller.Init(); err != nil {
			panic(err)
		}
		// 3. 初始化Api
		if err := ioc.Api.Init(); err != nil {
			panic(err)
		}

		if err := conf.C().Application.Start(); err != nil {
			panic(err)
		}
	},
}

func init() {
	Cmd.Flags().StringVarP(&testParam, "test", "t", "test", "test flag")
	// Root--> Init
	// -config xxxx
}
