package sharedtesting

import "github.com/stretchr/testify/mock"

// RendererMock …
type RendererMock struct {
	mock.Mock
}

// RenderComponents …
func (r *RendererMock) RenderComponents(output []byte) {
	r.Called(output)
}

// RenderLogs …
func (r *RendererMock) RenderLogs(output []byte) {
	r.Called(output)
}
