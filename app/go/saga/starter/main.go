package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {

	// Para iniciar debes abrir un objeto client. Este es un objeto pesado
	// Solo debes crear uno por proceso.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("No es posible crear cliente", err)
	}
	defer c.Close()

	// Configura las opciones para el workflow
	// ID:
	// Definimos identificador para el workflow
	//
	// Definimos el nombre del TaskQueue.
	// Qué es TaskQueue?:
	// 		Cuando un workflow invoca una actividad,
	// 		se envía el comando ScheduleActivityTask al servicio de Temporal.
	// 		Como resultado, el servicio actualiza el estado del workflow y
	//		envía una tarea de actividad a un worker que implementa la actividad.
	//		En lugar de llamar al worker directamente, se utiliza una cola intermedia.
	// 		Entonces, el servicio agrega una tarea de actividad a esta cola y un worker
	// 		recibe la tarea mediante una solicitud de encuesta larga.
	// 		Temporal llama a esta cola que se utiliza para distribuir tareas de actividad
	// 		en una cola de tareas de actividad.
	workflowOptions := client.StartWorkflowOptions{
		ID:        "viaje_workflowID", // identificador del workflow
		TaskQueue: "viaje",            // enrutar a taskqueue del worker donde esta el workflow definido en este caso es viaje.
	}

	// Ejecutamos el workflow.
	// En este caso no pasamos parámetros. Por simplicidad por ahora asumimos valores dentro del código y no para ser pasados como parámetros.
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, "viaje.reservar")

	// si falla la ejecución
	if err != nil {
		log.Fatalln("No es posible ejecutar workflow", err)
	}

	// logueamos el Id del workflow y el Id de su ejecución
	log.Println("Workflow iniciado", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Para este caso esperamos sincronamente (también es posible asincrono) a que el workflow se complete.
	var result []int
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("No es posible obtener resultado del workflow", err)
	}

	log.Println("Resultado:", result) // Deberia ser una secuencia de números con los Ids de transacción id[0] es del hotel, id[1] es del vehiculo, id[2] es del vuelo
}
