package business

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type UnitTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}

func (s *UnitTestSuite) Test_Workflow() {
	env := s.NewTestWorkflowEnvironment()

	// Realizar la prueba para la actividad individualmente
	// En el caso de este ejemplo la actividad es la función Activity
	env.OnActivity(Activity, "Sistecma").Return("Hola Sistecma", nil)

	// Realizar la prueba del workflow
	// En el caso de este ejemplo el workflow es la función Workflow
	env.ExecuteWorkflow(Workflow)

	// Validamos que el workflow este completo y que no hayan errores
	s.True(env.IsWorkflowCompleted())
	s.NoError(env.GetWorkflowError())

	env.AssertExpectations(s.T())
}
