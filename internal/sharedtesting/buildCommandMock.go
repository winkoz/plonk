package sharedtesting

import (
	"github.com/stretchr/testify/mock"
)

// BuildCommandMock …
type BuildCommandMock struct {
	mock.Mock
}

// Build …
func (oc *BuildCommandMock) Build(namespace string, isLatest bool) error {
	args := oc.Called(namespace, isLatest)
	return args.Error(0)
}

// Push …
func (oc *BuildCommandMock) Push(namespace string) error {
	args := oc.Called(namespace)
	return args.Error(0)
}
