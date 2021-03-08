package sharedtesting

import (
	"github.com/stretchr/testify/mock"
)

// VersionControllerCommandMock …
type VersionControllerCommandMock struct {
	mock.Mock
}

// Head …
func (oc *VersionControllerCommandMock) Head() (string, error) {
	args := oc.Called()
	return args.String(0), args.Error(1)
}
