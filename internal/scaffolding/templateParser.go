package scaffolding

import (
	"bytes"
	"text/template"

	"github.com/winkoz/plonk/internal/io/log"
)

// TemplateParser parses a template file with the appropiate values.
type TemplateParser interface {
	Parse(variables map[string]string, templateContent string) (string, error)
}

type templateParser struct{}

func (t templateParser) Parse(variables map[string]string, templateContent string) (string, error) {
	template, err := template.New("memory_template").Parse(templateContent)
	if err != nil {
		log.Errorf("Unable to parse template. %v", err)
		return "", err
	}

	buf := &bytes.Buffer{}
	err = template.Execute(buf, variables)
	if err != nil {
		log.Errorf("Unable to replace variables on template. %v", err)
		return "", err
	}

	return buf.String(), nil
}
