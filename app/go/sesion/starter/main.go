package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"

	"go.temporal.io/sdk/client"

	user "sesion/user"
)

func main() {
	var workflowID, input string
	flag.StringVar(&workflowID, "w", "workflow", "WorkflowID.")
	flag.StringVar(&input, "i", "{}", "Workflow input parameters.")
	flag.Parse()

	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	var we client.WorkflowRun
	var weError error

	var userState user.UserState
	if err := json.Unmarshal([]byte(input), &userState); err != nil {
		log.Fatalln("Unable to unmarshal workflow input parameters", err)
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "sesion",
	}

	we, weError = c.ExecuteWorkflow(context.Background(), workflowOptions, user.UserWorkflow, userState)

	if weError != nil {
		log.Fatalln("Unable to execute workflow", err)
	} else {
		log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
	}
}
