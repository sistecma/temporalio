package main

import (
	"context"
	"log"

	biz "child/biz"
	"flag"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	name := flag.String("name", "Sistecma", "a name")
	flag.Parse()

	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// This workflow ID can be user business logic identifier as well.
	workflowID := "parent-workflow_" + uuid.New()
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "child-workflow",
	}

	workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, biz.GreetingWorkflow, *name)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow",
		"WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = workflowRun.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Failure getting workflow result", err)
	}
	log.Println("Workflow result: ", result)
}
