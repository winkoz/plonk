package io

import (
	"github.com/stretchr/testify/mock"
)

// DuplicatorMock is a mock of Duplicator
type DuplicatorMock struct {
	mock.Mock
}

// CopyMultiple is a mock of CopyMultiple
func (d *DuplicatorMock) CopyMultiple(targetPath string, sourcePaths []FileLocation, transformator Transformator) error {
	args := d.Called(targetPath, sourcePaths)
	for _, source := range sourcePaths {
		transformator([]byte(source.ResolvedFilePath))
	}
	return args.Error(0)
}
