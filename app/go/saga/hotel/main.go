package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// Servicio / Actividad / Lógica de negocio
// En este caso la idea es simular la reserva de un hotel
func reservar(
	ctx context.Context,
	nombre string) (int, error) {

	logger := activity.GetLogger(ctx)
	logger.Info("Reservar hotel con nombre: " + nombre) // para este caso consideramos el print como transacción exitosa
	return rand.Int(), nil                              // número aleatorio que representa id de transacción ejecutada con éxito
}

// Servicio / Actividad / Lógica de negocio
// En este caso la idea es simular el reverso a una transacción de reserva de un hotel
func cancelar(
	ctx context.Context,
	nombre string,
	id int) (int, error) {

	logger := activity.GetLogger(ctx)
	logger.Info("Cancelar hotel con nombre: " + nombre + " id: " + strconv.Itoa(id)) // para este caso consideramos el print como transacción exitosa
	return rand.Int(), nil                                                           // número aleatorio que representa id de transacción ejecutada con éxito
}

func main() {
	// Para iniciar debes abrir un objeto client. Este es un objeto pesado
	// Solo debes crear uno por proceso.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("No es posible crear el cliente", err)
	}
	defer c.Close() // cierra el cliente al finalizar

	// creo el worker
	// como es el microservicio de hotel, tendrá un taskqueue llamado hotel

	// El taskqueue es importante porque nos permite enrutar las tareas a los workers

	w := worker.New(c, "hotel", worker.Options{}) // TaskQueue= hotel

	// registramos el identificador para la acción de reserva de hotel
	oa1 := activity.RegisterOptions{
		Name: "hotel.reservar", // identificador del servicio / actividad. Puede ser cualquier string que sea único para la aplicación
	}

	// asociamos el identificador con la función reservar
	w.RegisterActivityWithOptions(reservar, oa1) // registro el servicio / actividad con el worker

	// registramos el identificador para la acción de cancelar hotel
	oa2 := activity.RegisterOptions{
		Name: "hotel.cancelar", // identificador del servicio / actividad. Puede ser cualquier string que sea único para la aplicación
	}

	// asociamos el identificador con la función cancelar
	w.RegisterActivityWithOptions(cancelar, oa2) // registro el servicio / actividad con el worker

	// Ejecutar worker.
	// Este es un proceso demonio.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("No es posible ejecutar worker", err)
	}
}
