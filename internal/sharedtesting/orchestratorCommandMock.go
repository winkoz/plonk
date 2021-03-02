package sharedtesting

import (
	"github.com/stretchr/testify/mock"
)

// OrchestratorCommandMock …
type OrchestratorCommandMock struct {
	mock.Mock
}

// Deploy …
func (oc *OrchestratorCommandMock) Deploy(manifestPath string) error {
	args := oc.Called(manifestPath)
	return args.Error(0)
}

// Destroy …
func (oc *OrchestratorCommandMock) Destroy(env string) error {
	args := oc.Called(env)
	return args.Error(0)
}

// Diff …
func (oc *OrchestratorCommandMock) Diff(manifestPath string) error {
	args := oc.Called(manifestPath)
	return args.Error(0)
}

// Show …
func (oc *OrchestratorCommandMock) Show(env string) error {
	args := oc.Called(env)
	return args.Error(0)
}

// GetPods …
func (oc *OrchestratorCommandMock) GetPods(namespace string) ([]byte, error) {
	args := oc.Called(namespace)
	return args.Get(0).([]byte), args.Error(1)
}

// GetLogs …
func (oc *OrchestratorCommandMock) GetLogs(namespace string, component *string) ([]byte, error) {
	args := oc.Called(namespace, component)
	return args.Get(0).([]byte), args.Error(1)
}
