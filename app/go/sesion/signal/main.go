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
	var workflowID, signal string
	flag.StringVar(&workflowID, "w", "workflow", "WorkflowID.")
	flag.StringVar(&signal, "s", `{}`, "Signal data.")
	flag.Parse()

	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	var Event user.UserEvent
	if err := json.Unmarshal([]byte(signal), &Event); err != nil {
		log.Fatalln("Unable to unmarshal signal input parameters", err)
	}

	err = c.SignalWorkflow(context.Background(), workflowID, "", user.SignalName, Event)
	if err != nil {
		log.Fatalln("Unable to signal workflow", err)
	}
}
