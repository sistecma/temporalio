package workflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// counter var is the name of a variable that we want to observe from outside the workflow.
// it measures how many time we send a signal within a time interval
var counter int

// Check is a workflow that works in the following way:
// It has a timer that allows to specify a timeout for the workflow.
// The timeout is only resetted when a signal is sent to the workflow.
// If the timeout is reached the workflow is cancelled and counter is zeroed
// The workflow allows to query how many times a signal was sent within the timeout interval
func Check(ctx workflow.Context, signalName string, timeoutD time.Duration) error {

	var signalVal string
	var timeout bool = false

	// Register query handler to return count of signals within the time interval
	err := workflow.SetQueryHandler(ctx, "counter", func(input []byte) (int, error) {
		return counter, nil
	})

	// Validate in case of errors
	if err != nil {
		workflow.GetLogger(ctx).Info("Error executing query! ")
		return err
	}

	// Create a child context to link the timer
	childCtx, cancelHandler := workflow.WithCancel(ctx)

	// Create the signal channel
	signalChan := workflow.GetSignalChannel(ctx, signalName)

	// Define a selector to address timer and the signal input
	selector := workflow.NewSelector(ctx)

	// Use timer future to configure timeout feature, use childCtx
	timerFuture := workflow.NewTimer(childCtx, timeoutD)
	selector.AddFuture(timerFuture, func(f workflow.Future) {
		workflow.GetLogger(childCtx).Info("timed out! ")
		timeout = true
	})

	// Configure signal receiving feature
	selector.AddReceive(signalChan, func(channel workflow.ReceiveChannel, more bool) {
		channel.Receive(ctx, &signalVal)
		workflow.GetLogger(ctx).Info("Received signal! ", signalName, signalVal)
		counter++
		cancelHandler() // cancel the timer
	})

	workflow.GetLogger(ctx).Info("Waiting for timeout or signal! ")

	// Wait for the timeout to be reached or the signal to be received
	selector.Select(ctx)

	// This is just a business validation. We want to verify that we received a value and it is
	// different from "SOME_VALUE"
	if len(signalVal) > 0 && signalVal != "SOME_VALUE" {
		workflow.GetLogger(ctx).Info("signal value ok")
	} else if timeout { // If timeout is reached we zeroed the counter and return.
		counter = 0
		return nil
	}

	// At this point. The operation atomically completes the current execution and starts
	// a new execution of the Workflow with the same Workflow Id
	return workflow.NewContinueAsNewError(ctx, "timer.check", signalName, timeoutD)
}
