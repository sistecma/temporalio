package workflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

var counter int

func Check(ctx workflow.Context, signalName string, processingTimeThreshold time.Duration) error {

	var signalVal string
	var timeout bool = false

	// Register query handler to return count
	err := workflow.SetQueryHandler(ctx, "counter", func(input []byte) (int, error) {
		return counter, nil
	})

	if err != nil {
		workflow.GetLogger(ctx).Info("Error executing query! ")
		return err
	}

	childCtx, cancelHandler := workflow.WithCancel(ctx)
	signalChan := workflow.GetSignalChannel(ctx, signalName)

	selector := workflow.NewSelector(ctx)

	// use timer future to send notification email if processing takes too long
	timerFuture := workflow.NewTimer(childCtx, processingTimeThreshold)
	selector.AddFuture(timerFuture, func(f workflow.Future) {
		workflow.GetLogger(childCtx).Info("timed out! ")
		timeout = true
	})

	selector.AddReceive(signalChan, func(channel workflow.ReceiveChannel, more bool) {
		channel.Receive(ctx, &signalVal)
		workflow.GetLogger(ctx).Info("Received signal! ", signalName, signalVal)
		counter++
		cancelHandler()
	})

	workflow.GetLogger(ctx).Info("Waiting for timeout or signal! ")
	selector.Select(ctx)

	if len(signalVal) > 0 && signalVal != "SOME_VALUE" {
		workflow.GetLogger(ctx).Info("signal value ok")
	} else if timeout {
		counter = 0
		return nil
	}

	return workflow.NewContinueAsNewError(ctx, "timer.check", signalName, processingTimeThreshold)
}
