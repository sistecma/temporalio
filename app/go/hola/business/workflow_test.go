package business

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock para la actividad Activity del archivo hello.go
	env.OnActivity(Activity, mock.Anything, "Sistecma").Return("Hola Sistecma", nil)

	// Realizar la prueba del workflow
	// En el caso de este ejemplo el workflow es la función Workflow del archivo hello.go
	env.ExecuteWorkflow(Workflow, "Sistecma")

	// Validamos que el workflow este completo y que no hayan errores
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	// Obtenemos y Evaluamos el resultado esperado del workflow
	var result string
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, "Hola Sistecma", result)
}
