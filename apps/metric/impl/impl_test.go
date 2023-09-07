package impl_test

import (
	"context"

	"github.com/solodba/dbexporter/apps/metric"
	"github.com/solodba/dbexporter/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc metric.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(metric.AppName).(metric.Service)
}
