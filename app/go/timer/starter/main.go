package main

import (
	"context"
	"log"
	"time"

	workflowf "timer/workflow"

	"go.temporal.io/sdk/client"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "timer_hmo",
		TaskQueue: "timer",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflowf.Check, "mysignal", time.Second*10)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
