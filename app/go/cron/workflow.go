package cron

import (
	"go.temporal.io/sdk/workflow"
)

// CronWorkflow is the sample cron workflow.
func CronWorkflow(ctx workflow.Context) (error) {
	workflow.GetLogger(ctx).Info("Cron workflow started.", "StartTime", workflow.Now(ctx))
	return  nil
}
