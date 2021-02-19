package main

import (
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// workflow que implementa el patrón SAGA y la orquestación de los microservicios
// retorna los ids resultados de cada microservicio.
// La secuencia de números con los Ids de transacción id[0] es del hotel, id[1] es del vehiculo, id[2] es del vuelo
func reservar(ctx workflow.Context) ([]int, error) {

	// definimos un logger para el workflow
	logger := workflow.GetLogger(ctx)

	// definimos una opción de actividad general (por facilidad y simplicidad)
	// las opciones de actividad permiten configurar granularmente las actividades desde el punto de vista de Temporal
	// timeouts cortos en duración por propósitos del ejercicio
	aa := workflow.ActivityOptions{
		// tiempo máximo que puede transcurrir desde que un workflow solicite la ejecución de una actividad
		// hasta que un worker inicie la ejecución de dicha actividad.
		// Si se dispara este timeout es indicativo de que el/los workers que registran la actividad estan:
		// o abajo o no pueden mantener la velocidad de despacho de tareas.
		ScheduleToStartTimeout: time.Second * 3,

		// tiempo máximo dentro del cual se puede ejecutar una tarea una vez que es tomada por un worker.
		StartToCloseTimeout: time.Second * 3,
	}

	// inicializamos el contexto genérico y cargamos la opción de actividad general
	ctx = workflow.WithActivityOptions(ctx, aa)

	// INICIO DE IMPLEMENTACIÓN DE SAGA ----------------------------------------

	// inicializamos los resultados de los ids de cada microservicio
	result := make([]int, 3)
	result[0] = -1 // microservicio de hotel
	result[1] = -1 // microservicio de vehiculo
	result[2] = -1 // microservicio de vuelo

	// invocar microservicio de hotel
	var idHotel int
	var nombreHotel string = "Royal Cool Hotel"                          // definimos para el ejemplo el nombre de un hotel cualquiera
	var ctxHotel workflow.Context = workflow.WithTaskQueue(ctx, "hotel") // indicamos que se enrute a la cola del microservicio de hotel

	errHotel := workflow.ExecuteActivity(ctxHotel, "hotel.reservar", nombreHotel).Get(ctxHotel, &idHotel) // indicamos que actividad ejecutar

	if errHotel != nil {
		logger.Error("Falla ejecutando la actividad hotel.reservar ", "Error", errHotel)
		return result, errHotel
	}
	result[0] = idHotel // asignamos al index 0 del result (que retornará el resultado)

	// invocar microservicio de vehiculo
	var idVehiculo int
	var nombreVehiculo string = "Cool Taxis"                                   // definimos para el ejemplo el nombre de un vehiculo cualquiera
	var ctxVehiculo workflow.Context = workflow.WithTaskQueue(ctx, "vehiculo") // indicamos que se enrute a la cola del microservicio de vehiculo

	errVehiculo := workflow.ExecuteActivity(ctxVehiculo, "vehiculo.reservar", nombreVehiculo).Get(ctxVehiculo, &idVehiculo) // indicamos que actividad ejecutar
	if errVehiculo != nil {
		logger.Error("Falla ejecutando la actividad vehiculo.reservar ", "Error", errVehiculo)

		// reverso de reserva de hotel
		var idHotelRev int
		// ojo aqui notar que usamos el ctxHotel porque queremos se enrute al microservicio hotel
		errHotelRev := workflow.ExecuteActivity(ctxHotel, "hotel.cancelar", nombreHotel, idHotel).Get(ctxHotel, &idHotelRev)
		if errHotelRev != nil {
			// reverso fallido. posiblemente es necesario hacer una actividad de conciliación manual.
			logger.Error("Falla ejecutando la actividad hotel.cancelar ", "Error", errHotelRev)
		}
		return result, errVehiculo
	}
	result[1] = idVehiculo

	// invocar microservicio de Vuelo
	var idVuelo int
	var nombreVuelo string = "Cool airlines"                             // definimos para el ejemplo el nombre de una aerolinea cualquiera
	var ctxVuelo workflow.Context = workflow.WithTaskQueue(ctx, "vuelo") // indicamos que se enrute a la cola del microservicio de vuelo

	errVuelo := workflow.ExecuteActivity(ctxVuelo, "vuelo.reservar", nombreVuelo).Get(ctxVuelo, &idVuelo) // indicamos que actividad ejecutar
	if errVuelo != nil {
		logger.Error("Falla ejecutando la actividad vuelo.reservar ", "Error", errVuelo)

		// reverso de reserva de vehiculo
		var idVehiculoRev int
		// ojo aqui notar que usamos el ctxVehiculo porque queremos se enrute al microservicio vehiculo
		errVehiculoRev := workflow.ExecuteActivity(ctxVehiculo, "vehiculo.cancelar", nombreVehiculo, idVehiculo).Get(ctxVehiculo, &idVehiculoRev)
		if errVehiculoRev != nil {
			// reverso fallido. posiblemente es necesario hacer una actividad de conciliación manual.
			logger.Error("Falla ejecutando la actividad vehiculo.cancelar ", "Error", errVehiculoRev)
		}

		// reverso de reserva de hotel
		var idHotelRev int
		// ojo aqui notar que usamos el ctxHotel porque queremos se enrute al microservicio hotel
		errHotelRev := workflow.ExecuteActivity(ctxHotel, "hotel.cancelar", nombreHotel, idHotel).Get(ctxHotel, &idHotelRev)
		if errHotelRev != nil {
			// reverso fallido. posiblemente es necesario hacer una actividad de conciliación manual.
			logger.Error("Falla ejecutando la actividad hotel.cancelar ", "Error", errHotelRev)
		}

		return result, errVuelo
	}
	result[2] = idVuelo

	// FIN DE IMPLEMENTACIÓN DE SAGA -------------------------------------------

	return result, nil
}

func main() {
	// Para iniciar debes abrir un objeto client. Este es un objeto pesado
	// Solo debes crear uno por proceso.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("No es posible crear el cliente", err)
	}
	defer c.Close()

	// creo el worker
	w := worker.New(c, "viaje", worker.Options{}) // TaskQueue= viaje

	ow1 := workflow.RegisterOptions{
		Name: "viaje.reservar",
	}
	w.RegisterWorkflowWithOptions(reservar, ow1)

	// Ejecutar worker.
	// Este es un proceso demonio.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("No es posible ejecutar worker", err)
	}

}
