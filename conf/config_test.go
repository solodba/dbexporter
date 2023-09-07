package conf_test

import (
	"testing"

	"github.com/solodba/dbexporter/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("test/test.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().App)
	t.Log(conf.C().MySQL)
	t.Log(conf.C().MySQL.GetDb())
}

func TestLoadConfigFromEnv(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().App)
	t.Log(conf.C().MySQL.GetDb())
}
