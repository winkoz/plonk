package sharedtesting

import (
	"github.com/stretchr/testify/mock"
)

// ExecutorMock …
type ExecutorMock struct {
	mock.Mock
}

// Run …
func (e *ExecutorMock) Run(command string, arg ...string) ([]byte, error) {
	args := e.Called(command, arg)
	return args.Get(0).([]byte), args.Error(1)
}
