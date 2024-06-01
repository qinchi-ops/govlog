package conf_test

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/qinchi-ops/govlog/vblog/conf"
)

func TestToYaml(t *testing.T) {
	t.Log(conf.Default().ToYaml())
}

func TestLoadFromEnv(t *testing.T) {
	os.Setenv("DATASOURCE_HOST", "env test")
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().ToYaml())
	assert.Equal(t, conf.C().Mysql.Username, "env test")
}

func TestLoadFromYam(t *testing.T) {
	err := conf.LoadConfigFromYaml("./application.yaml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().ToYaml())
}
