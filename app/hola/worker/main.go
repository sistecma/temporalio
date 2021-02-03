package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	business "hola/business"
)

func main() {
	// Para iniciar debes abrir un objeto client. Este es un objeto pesado
	// Solo debes crear uno por proceso.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("No es posible crear el cliente", err)
	}
	defer c.Close()

	w := worker.New(c, "hello", worker.Options{}) // el mismo nombre del TaskQueue

	// Para este caso registro ambos Workflow y Actividad.
	// Es posible registrarlo por separado si es requerido.
	w.RegisterWorkflow(business.Workflow)
	w.RegisterActivity(business.Activity)

	// Ejecutar worker.
	// Este es un proceso demonio.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("No es posible ejecutar worker", err)
	}
}
