package metrics

// MysqlResponse结构体
type MysqlResponse struct {
	Name  string
	Value string
}

// MysqlResponse结构体构造函数
func NewMysqlResponse() *MysqlResponse {
	return &MysqlResponse{}
}
