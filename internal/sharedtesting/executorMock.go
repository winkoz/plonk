package sharedtesting

import (
	"github.com/stretchr/testify/mock"
)

// ExecutorMock …
type ExecutorMock struct {
	mock.Mock
}

// Run …
func (e *ExecutorMock) Run(command string, arg ...string) error {
	args := e.Called(command, arg)
	return args.Error(0)
}
