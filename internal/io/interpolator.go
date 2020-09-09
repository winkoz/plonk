package io

import (
	"crypto/sha1"
	"encoding/hex"
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

// SubstituteValues replaces all instances of 'key' with its respective 'value' from the `source` map in the `template` string and returns the applied template `string`.
func (r interpolator) SubstituteValues(source map[string]string, template string) (string, error) {
	result := template
	hashedMap := map[string]string{}
	for key, value := range source {
		hasher := sha1.New()
		hasher.Write([]byte(key))
		hashedKey := hex.EncodeToString(hasher.Sum(nil))
		result = strings.ReplaceAll(result, fmt.Sprintf("%s%s", InterpolatorSignaler, key), fmt.Sprintf("%s{%s}", InterpolatorSignaler, hashedKey))
		hashedMap[hashedKey] = value
	}
	for key, value := range hashedMap {
		result = strings.ReplaceAll(result, fmt.Sprintf("%s{%s}", InterpolatorSignaler, key), value)
	}
	return result, nil
}
