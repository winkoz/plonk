package scaffolding

import "github.com/stretchr/testify/mock"

type TemplateReaderMock struct {
	mock.Mock
}

func (trm *TemplateReaderMock) Read(templateName string) (TemplateData, error) {
	args := trm.Called(templateName)
	return args.Get(0).(TemplateData), args.Error(1)
}
