package impl

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/solodba/dbexporter/apps/collector"
	"github.com/solodba/dbexporter/apps/metrics"
	"github.com/solodba/mcube/apps"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	svc                  metrics.Service
	threadsConnectedDesc *prometheus.Desc
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return collector.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	i.svc = apps.GetInternalApp(metrics.AppName).(metrics.Service)
	return nil
}

func newThreadsConnectedDesc() *prometheus.Desc {
	return prometheus.NewDesc(
		"threads_connected",
		"The number of mysql threads connected.",
		[]string{},
		prometheus.Labels{"instance_name": "master"},
	)
}

// 注册实例类
func init() {
	apps.RegistryInternalApp(svc)
	svc.threadsConnectedDesc = newThreadsConnectedDesc()
	prometheus.MustRegister(svc)
}
