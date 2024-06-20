package test

import (
	"github.com/qinchi-ops/govlog/vblog/conf"
	"github.com/qinchi-ops/govlog/vblog/ioc"
	"github.com/spf13/cobra"

	// 注册所有的业务模块
	_ "github.com/qinchi-ops/govlog/vblog/apps"
)

func DevelopmentSetup() {
	//加载配置，单元测试通过环境变量读取，vscode传递进来
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	// // 2. 初始化Ioc
	// if err := ioc.Controller.Init(); err != nil {
	// 	panic(err)
	// }
	// // 3. 初始化Api
	// if err := ioc.Api.Init(); err != nil {
	// 	panic(err)
	// }

	// 2. 初始化Ioc
	cobra.CheckErr(ioc.Controller.Init())
	// 3. 初始化Api
	cobra.CheckErr(ioc.Api.Init())

}
