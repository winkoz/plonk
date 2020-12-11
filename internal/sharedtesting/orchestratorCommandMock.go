package sharedtesting

import (
	"github.com/stretchr/testify/mock"
)

// OrchestratorCommandMock …
type OrchestratorCommandMock struct {
	mock.Mock
}

// Deploy …
func (oc *OrchestratorCommandMock) Deploy(env string, manifestPath string) error {
	args := oc.Called(env, manifestPath)
	return args.Error(0)
}

// Diff …
func (oc *OrchestratorCommandMock) Diff(env string, manifestPath string) error {
	args := oc.Called(env, manifestPath)
	return args.Error(0)
}

// Show …
func (oc *OrchestratorCommandMock) Show(env string) error {
	args := oc.Called(env)
	return args.Error(0)
}
