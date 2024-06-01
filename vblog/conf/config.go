package conf

import "gopkg.in/yaml.v3"

func Default() *Config {
	return &Config{
		Application: &application{
			Host: "127.0.0.1",
			Port: 8080,
		},
		Mysql: &mySQL{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "go15",
			Username: "",
			Password: "",
			Debug:    true,
		},
	}
}

// 这个对象就是程序配置
type Config struct {
	Application *application `json:"app" yaml:"app" toml:"app"`
	Mysql       *mySQL       `json:"mysql" yaml:"mysql" toml:"mysql"`
}

func (c *Config) ToYaml() string {
	out, _ := yaml.Marshal(c)
	return string(out)
}

// 应用服务
type application struct {
	Host string `toml:"host" json:"host" yaml:"host"`
	Port int    `toml:"port" json:"port" yaml:"port"`
}

type mySQL struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DATABASE"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`
}
