package main

import (
	"github.com/qinchi-ops/govlog/vblog/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
