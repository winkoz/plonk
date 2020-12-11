package sharedtesting

import (
	"github.com/stretchr/testify/mock"
	"github.com/winkoz/plonk/internal/io"
)

// SecretReaderMock …
type SecretReaderMock struct {
	mock.Mock
}

// GetSecretsFromFile …
func (sr *SecretReaderMock) GetSecretsFromFile(projectName string, env string) (result io.DeploySecrets, err error) {
	args := sr.Called(projectName, env)
	return args.Get(0).(io.DeploySecrets), args.Error(1)
}
