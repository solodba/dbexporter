package impl

import (
	"context"
	"strconv"

	"gitee.com/code-horse-mi/mcube/logger"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ctx = context.Background()
)

func (i *impl) Describe(ch chan<- *prometheus.Desc) {
	ch <- i.threadsConnectedDesc
}

func (i *impl) Collect(ch chan<- prometheus.Metric) {
	mysqlResponse, err := i.svc.GetMysqlThreadsConnected(ctx)
	if err != nil {
		logger.L().Error().Msgf("get mysql threads connected error: %s", err.Error())
		return
	}
	threadConnected, err := strconv.ParseFloat(mysqlResponse.Value, 64)
	if err != nil {
		logger.L().Error().Msgf("get mysql threads connected error: %s", err.Error())
		return
	}
	ch <- prometheus.MustNewConstMetric(i.threadsConnectedDesc, prometheus.GaugeValue, threadConnected)
}
