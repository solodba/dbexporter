package collector

import "github.com/prometheus/client_golang/prometheus"

// 模块名称
const (
	AppName = "collector"
)

// 收集器相关服务接口
type Service interface {
	// prometheus收集接口
	prometheus.Collector
}
