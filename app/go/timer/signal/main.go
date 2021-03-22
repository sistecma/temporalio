package main

import (
	"context"
	"log"

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

	err = c.SignalWorkflow(context.Background(), "timer_hmo", "", "mysignal", "REAL_VALUE")
	if err != nil {
		log.Fatalln("Unable to signal workflow", err)
	}
}
