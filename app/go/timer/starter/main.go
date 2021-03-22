package main

import (
	"context"
	"log"
	"time"

	workflowf "timer/workflow"

	"go.temporal.io/sdk/client"
)

// Starter kickstars the workflow
func main() {
	// The client is a heavyweight object that should be created once per process.
	// Make the connection with Temporal
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})

	// Validate if there is any error
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Set workflow id and queue names
	workflowOptions := client.StartWorkflowOptions{
		ID:        "timer_sistecma",
		TaskQueue: "timer",
	}

	// Invoke the workflow
	// Set signal name and timeout duration
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflowf.Check, "mysignal", time.Second*10)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
