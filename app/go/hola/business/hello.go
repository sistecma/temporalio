package business

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

// La función Workflow es la definición del workflow.
// Para este caso solo maneja la configuración del llamado de la actividad
func Workflow(ctx workflow.Context, name string) (string, error) {

	// Configuramos el comportamiento para las actividades desde la función Workflow
	// En Temporal todo es manejado asíncronamente.
	// Las actividades tienen configuraciones de timeouts a ser considerado
	//
	// ScheduleToStartTimeout:
	//    Tiempo que puede esperar una tarea para que un worker de actividad
	//    la recoja después de que un flujo de trabajo la programe.
	//    Si no hay workers disponibles para procesar esta tarea durante
	//    la duración especificada, la tarea expirará.
	//
	// StartToCloseTimeout:
	//    Tiempo que puede tardar una tarea en completarse
	//    después de que un flujo de trabajo la programe.
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute, // minuto
		StartToCloseTimeout:    time.Minute, // minuto
	}
	ctx = workflow.WithActivityOptions(ctx, ao) // configura las opciones de las actividades

	// Como logger siempre usa el que es provisto por Temporal dentro de un workflow.
	logger := workflow.GetLogger(ctx)

	logger.Info("Workflow Hello ha iniciado", "name", name)

	var result string

	// Ejecución de la actividad dentro del workflow.
	// Recuerda que el workflow preserva el estado y para ejecutar lógica de
	// negocio usa a las actividades.
	err := workflow.ExecuteActivity(ctx, Activity, name).Get(ctx, &result)

	// si existe error
	if err != nil {
		logger.Error("Actividad fallida.", "Error", err)
		return "", err
	}

	logger.Info("Workflow Hello se completó.", "result", result)

	return result, nil
}

// Lógica de negocio. En este caso es una función que recibe un nombre
// Y retorna Hola + el contenido del nombre
func Activity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)

	// retorna Hola + name
	return "Hola " + name, nil
}
