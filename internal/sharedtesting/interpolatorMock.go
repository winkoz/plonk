package sharedtesting

import "github.com/stretchr/testify/mock"

// InterpolatorMock is a mock of Interpolator
type InterpolatorMock struct {
	mock.Mock
}

// SubstituteValues mocks SubstituteValues
func (r *InterpolatorMock) SubstituteValues(source map[string]string, target string) (string, error) {
	args := r.Called(source, target)
	return args.String(0), args.Error(1)
}
