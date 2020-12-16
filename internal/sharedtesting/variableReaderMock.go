package sharedtesting

import (
	"github.com/stretchr/testify/mock"
	"github.com/winkoz/plonk/internal/io"
)

// VariableReaderMock …
type VariableReaderMock struct {
	mock.Mock
}

// GetVariablesFromFile …
func (vr *VariableReaderMock) GetVariablesFromFile(projectName string, env string) (io.DeployVariables, error) {
	args := vr.Called(projectName, env)
	return args.Get(0).(io.DeployVariables), args.Error(1)
}
