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
