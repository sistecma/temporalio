package main

import (
	"go.temporal.io/sdk/workflow"
)

// ParentWorkflow. This sample workflow demonstrates how to use invoke child workflow from parent workflow execution.
// When a child workflow execution is starting a new run the parent execution is notified only after the completion of last run.
// In this case this workflow receives a message as string and forward to the child workflow
func ParentWorkflow(ctx workflow.Context, parent string) (string, error) {
	logger := workflow.GetLogger(ctx)
	// use options for ChildWorkflow
	cwo := workflow.ChildWorkflowOptions{}
	// set the context with child options
	ctx = workflow.WithChildOptions(ctx, cwo)
	// execute child workflow
	var result string
	err := workflow.ExecuteChildWorkflow(ctx, ChildWorkflow, parent).Get(ctx, &result)
	if err != nil {
		logger.Error("Parent execution received child execution failure.", "Error", err)
		return "", err
	}
	logger.Info("Parent execution completed.", "Result from child: ", result)
	return result, nil
}

// ChildWorkflow. In this case is defined in the same go file and deployed in the same workfer but,
// It could be in a another worker, cloud or whatever as long as Temporal service reaches it,
// then it can be executed. In this case it just sends back the same message but adding a comment
func ChildWorkflow(ctx workflow.Context, child string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Childworkflow execution.", "Message is: ", child)

	return "Childworkflow executed, original message is: " + child, nil
}
