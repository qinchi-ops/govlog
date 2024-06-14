package main

import (
	"github.com/qinchi-ops/govlog/vblog/cmd"

	//业务对象注册
	_ "github.com/qinchi-ops/govlog/vblog/apps"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
