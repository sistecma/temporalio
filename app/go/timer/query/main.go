package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

// Query how many times we send a signal within a time interval
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

	// Make a query. Point to workflow-id= timer_sistecma and query counter var
	resp, err := c.QueryWorkflow(context.Background(), "timer_sistecma", "", "counter")

	// Validate if there is any error
	if err != nil {
		log.Fatalln("Unable to query workflow", err)
	}

	// Declare result, get it, and validate if there is any error
	var result interface{}
	if err := resp.Get(&result); err != nil {
		log.Fatalln("Unable to decode query result", err)
	}

	// Print in console
	log.Println("Received query result", "Result", result)
}
