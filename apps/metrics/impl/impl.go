package impl

import (
	"database/sql"

	"github.com/solodba/dbexporter/apps/metrics"
	"github.com/solodba/dbexporter/conf"
	"github.com/solodba/mcube/apps"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	db *sql.DB
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return metrics.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	db, err := conf.C().MySQL.GetDb()
	if err != nil {
		return err
	}
	i.db = db
	return nil
}

// 注册实例类
func init() {
	apps.RegistryInternalApp(svc)
}
