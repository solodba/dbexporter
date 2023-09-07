package impl_test

import (
	"testing"

	"github.com/solodba/dbexporter/test/tools"
)

func TestGetMysqlThreadsConnected(t *testing.T) {
	res, err := svc.GetMysqlThreadsConnected(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(res))
}
