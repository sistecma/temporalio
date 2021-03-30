package biz

import (
	"go.temporal.io/sdk/workflow"
)

// GreetingWorkflow (parent) workflow definition
func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)

	cwo := workflow.ChildWorkflowOptions{}
	ctx = workflow.WithChildOptions(ctx, cwo)

	var result1 string
	err1 := workflow.ExecuteChildWorkflow(ctx, HelloWorkflow, name).Get(ctx, &result1)
	if err1 != nil {
		logger.Error("Parent execution received child execution failure.", "Error", err1)
		return "", err1
	}

	var result2 string
	err2 := workflow.ExecuteChildWorkflow(ctx, HowAreYouWorkflow, name).Get(ctx, &result2)
	if err2 != nil {
		logger.Error("Parent execution received child execution failure.", "Error", err2)
		return "", err2
	}

	result := result1 + ". " + result2
	logger.Info("Parent execution completed.", result)
	return result, nil
}

// Hello (child) workflow definition
func HelloWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	greeting := "Hello " + name
	logger.Info("Hello workflow execution: " + greeting)
	return greeting, nil
}

// How are you? (child) workflow definition
func HowAreYouWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	greeting := "How are you, " + name + "?"
	logger.Info("How are you workflow execution: " + greeting)
	return greeting, nil
}
