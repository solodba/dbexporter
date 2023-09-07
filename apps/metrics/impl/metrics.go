package impl

import (
	"context"

	"github.com/solodba/dbexporter/apps/metrics"
)

// 获取线程连接数
func (i *impl) GetMysqlThreadsConnected(ctx context.Context) (*metrics.MysqlResponse, error) {
	sql := `show global status like 'Threads_connected'`
	row := i.db.QueryRowContext(ctx, sql)
	res := metrics.NewMysqlResponse()
	err := row.Scan(&res.Name, &res.Value)
	if err != nil {
		return nil, err
	}
	return res, nil
}
