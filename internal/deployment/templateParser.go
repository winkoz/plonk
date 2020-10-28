package deployment

import (
	"bytes"
	"text/template"

	"github.com/winkoz/plonk/internal/io/log"
)

// TemplateParser parses a template file with the appropiate values.
type TemplateParser interface {
	Parse(variables map[string]interface{}, templateContent string) (string, error)
}

type templateParser struct{}

// NewTemplateParser returns a fully initialised TemplateParser
func NewTemplateParser() TemplateParser {
	return templateParser{}
}

func (t templateParser) Parse(variables map[string]interface{}, templateContent string) (result string, err error) {
	signal := log.StartTrace("Parse")
	defer log.StopTrace(signal, err)

	template, err := template.New("memory_template").Parse(templateContent)
	if err != nil {
		log.Errorf("Unable to parse template. %v", err)
		return result, err
	}

	buf := &bytes.Buffer{}
	err = template.Execute(buf, variables)
	if err != nil {
		log.Errorf("Unable to replace variables on template. %v", err)
		return result, err
	}

	result = buf.String()
	return result, nil
}
