package scaffolding

import "github.com/stretchr/testify/mock"

type templateReaderMock struct {
	mock.Mock
}

func (trm *templateReaderMock) Read(templateName string) (TemplateData, error) {
	args := trm.Called(templateName)
	return args.Get(0).(TemplateData), args.Error(1)
}
