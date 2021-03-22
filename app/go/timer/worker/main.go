package main

import (
	"log"

	workflowf "timer/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// Worker configuration
func main() {

	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("It is not possible to create the client", err)
	}
	defer c.Close()

	// Create the worker
	w := worker.New(c, "timer", worker.Options{}) // TaskQueue= timer

	ow1 := workflow.RegisterOptions{
		Name: "timer.check",
	}
	w.RegisterWorkflowWithOptions(workflowf.Check, ow1)

	// Execute the worker
	// This is a daemon process
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("It is not possible to run the worker", err)
	}

}
