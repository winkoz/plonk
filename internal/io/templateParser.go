package io

import (
	"bytes"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/winkoz/plonk/internal/io/log"
)

// TemplateParser parses a template file with the appropiate values.
type TemplateParser interface {
	Parse(variables map[string]interface{}, templateContent string) (string, error)
}

type templateParser struct {
	service         Service
	dataManipulator DataManipulator
}

// NewTemplateParser returns a fully initialised TemplateParser
func NewTemplateParser() TemplateParser {
	service := NewService()
	return templateParser{
		service:         service,
		dataManipulator: NewDataManipulator(service),
	}
}

func (t templateParser) Parse(variables map[string]interface{}, templateContent string) (result string, err error) {
	signal := log.StartTrace("Parse")
	defer log.StopTrace(signal, err)

	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title":          strings.Title,
		"readFile":       t.service.ReadFile,
		"readFileToStr":  t.service.ReadFileToString,
		"strToBytes":     t.dataManipulator.StringToBytes,
		"base64Encode":   t.dataManipulator.Base64Encode,
		"walkDirectory":  t.service.WalkDirectory,
		"baseFilename":   filepath.Base,
		"yamlArrayToObj": t.dataManipulator.YamlToMapArray,
		"indent":         t.dataManipulator.Indent,
	}

	template, err := template.
		New("memory_template").
		Funcs(funcMap).
		Parse(templateContent)
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
