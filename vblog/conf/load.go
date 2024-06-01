package conf

import (
	"os"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
)

// 配置加载
// file/env/...   --->Config
//全局一份

// config 全局变量，通过函数对外提供访问
var config *Config

func C() *Config {
	//没有配置文件怎么办
	//默认配置，方便开发者
	if config == nil {
		config = Default()
	}
	return config
}

// 加载配置,把外部配置读取到config全局变量来
// yaml 文件yaml--->Config
func LoadConfigFromYaml(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	config = C()

	return yaml.Unmarshal(content, config)
}

// 环境变量里面读取配置
func LoadConfigFromEnv() error {
	config = C()
	//MYSQL_DB <---> DB

	return env.Parse(config)
}
