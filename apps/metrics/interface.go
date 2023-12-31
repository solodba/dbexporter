package metrics

import "context"

// 模块名称
const (
	AppName = "metrics"
)

// 获取MySQL监控指标服务接口
type Service interface {
	// 获取线程连接数
	GetMysqlThreadsConnected(context.Context) (*MysqlResponse, error)
}
