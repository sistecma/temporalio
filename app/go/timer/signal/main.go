package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

// Send a signal to the workflow
func main() {

	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})

	// Validate in case of error
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Send signal "mysignal" with value "REAL_VALUE" to workflow id "timer_sistecma"
	err = c.SignalWorkflow(context.Background(), "timer_sistecma", "", "mysignal", "REAL_VALUE")
	if err != nil {
		log.Fatalln("Unable to signal workflow", err)
	}
}
