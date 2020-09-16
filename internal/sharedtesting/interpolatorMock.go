package sharedtesting

import "github.com/stretchr/testify/mock"

// InterpolatorMock is a mock of Interpolator
type InterpolatorMock struct {
	mock.Mock
}

// SubstituteValues mocks SubstituteValues
func (i *InterpolatorMock) SubstituteValues(source map[string]string, target string) (string, error) {
	args := i.Called(source, target)
	return args.String(0), args.Error(1)
}

// SubstituteValuesInMap …
func (i *InterpolatorMock) SubstituteValuesInMap(source map[string]string, target map[string]string) (map[string]string, error) {
	args := i.Called(source, target)
	return args.Get(0).(map[string]string), args.Error(1)
}
