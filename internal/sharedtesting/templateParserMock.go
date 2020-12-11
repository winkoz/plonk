package sharedtesting

import (
	"github.com/stretchr/testify/mock"
)

// TemplateParserMock …
type TemplateParserMock struct {
	mock.Mock
}

// Parse …
func (tp *TemplateParserMock) Parse(variables map[string]interface{}, templateContent string) (string, error) {
	args := tp.Called(variables, templateContent)
	return args.String(0), args.Error(1)
}
