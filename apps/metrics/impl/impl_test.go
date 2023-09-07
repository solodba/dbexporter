package impl_test

import (
	"context"

	"github.com/solodba/dbexporter/apps/metrics"
	"github.com/solodba/dbexporter/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc metrics.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(metrics.AppName).(metrics.Service)
}
