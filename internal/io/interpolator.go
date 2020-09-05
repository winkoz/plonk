package io

import (
	"fmt"
	"strings"
)

type interpolator struct{}

// InterpolatorSignaler indicates the start of a variable that should be sustitued in a template.
const InterpolatorSignaler string = "$"

// Interpolator manages variable substitution
type Interpolator interface {
	SubstituteValues(source map[string]string, template string) (string, error)
}

// NewInterpolator returns a fully initialised Interpolator
func NewInterpolator() Interpolator {
	return interpolator{}
}

func (r interpolator) SubstituteValues(source map[string]string, template string) (string, error) {
	result := template
	for key, value := range source {
		result = strings.ReplaceAll(result, fmt.Sprintf("%s%s", InterpolatorSignaler, key), value)
	}
	return result, nil
}
