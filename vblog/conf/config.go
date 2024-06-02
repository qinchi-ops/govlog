package conf

import (
	"fmt"
	"sync"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
			Username: "root",
			Password: "root",
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
	//保证他是单离
	db *gorm.DB
	//加锁
	lock sync.Mutex
}

// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (m *mySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
	)
}

func (m *mySQL) GetDB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.db == nil {
		db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.db = db
		//补充debug信息
		if m.Debug {
			m.db = db.Debug()
		}

	}
	return m.db
}
