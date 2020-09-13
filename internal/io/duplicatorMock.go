package io

import (
	"github.com/stretchr/testify/mock"
)

// DuplicatorMock is a mock of Duplicator
type DuplicatorMock struct {
	mock.Mock
}

// CopyMultiple is a mock of CopyMultiple
func (d *DuplicatorMock) CopyMultiple(sourcePath string, targetPath string, sources []string, transformator Transformator) error {
	args := d.Called(sourcePath, targetPath, sources)
	for _, source := range sources {
		transformator([]byte(source))
	}
	return args.Error(0)
}
