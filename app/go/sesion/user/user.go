package user

import (
	"go.temporal.io/sdk/workflow"
)

type (
	// UserState kept within workflow and passed from one run to another on ContinueAsNew
	UserState struct {
		Counter int
	}

	// TripEvent passed in as signal to TripWorkflow
	UserEvent struct {
		ID    string
		Total int
	}
)

const (
	// SignalName is the signal name for completion event
	SignalName = "user_event"

	// QueryName is the query type name
	QueryName = "counter"
)

func UserWorkflow(ctx workflow.Context, state UserState) error {
	logger := workflow.GetLogger(ctx)

	workflowID := workflow.GetInfo(ctx).WorkflowExecution.ID

	logger.Info("Workflow Started for User ", "User", workflowID, "Counter", state.Counter)

	// Register query handler to return count
	err := workflow.SetQueryHandler(ctx, QueryName, func(input []byte) (int, error) {
		return state.Counter, nil
	})

	if err != nil {
		logger.Info("SetQueryHandler failed.", "Error", err)
		return err
	}

	// Ch to wait on completed event signals
	logger.Info("before workflow.GetSignalChannel")
	Ch := workflow.GetSignalChannel(ctx, SignalName)
	logger.Info("after workflow.GetSignalChannel")
	for i := 0; i < 5; i++ { // wait for 5 signals
		var user UserEvent
		logger.Info("before Ch.Receive")
		Ch.Receive(ctx, &user)
		logger.Info("after Ch.Receive")
		logger.Info("User complete event received.", "ID", user.ID, "Total", user.Total)
		state.Counter++ // counter of signals
	}

	logger.Info("Starting a new run.", "Counter", state.Counter)
	return workflow.NewContinueAsNewError(ctx, "Workflow", state)
}
