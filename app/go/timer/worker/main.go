package main

import (
	"log"

	workflowf "timer/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func main() {
	// Para iniciar debes abrir un objeto client. Este es un objeto pesado
	// Solo debes crear uno por proceso.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("No es posible crear el cliente", err)
	}
	defer c.Close()

	// creo el worker
	w := worker.New(c, "timer", worker.Options{}) // TaskQueue= timer

	ow1 := workflow.RegisterOptions{
		Name: "timer.check",
	}
	w.RegisterWorkflowWithOptions(workflowf.Check, ow1)

	// Ejecutar worker.
	// Este es un proceso demonio.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("No es posible ejecutar worker", err)
	}

}
