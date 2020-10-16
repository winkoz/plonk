package io

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/winkoz/plonk/internal/io/log"
)

type interpolator struct{}

// InterpolatorSignaler indicates the start of a variable that should be sustitued in a template.
const InterpolatorSignaler string = "$"

// Interpolator manages variable substitution
type Interpolator interface {
	SubstituteValues(source map[string]string, target string) string
	SubstituteValuesInMap(source map[string]string, target map[string]string) map[string]string
}

// NewInterpolator returns a fully initialised Interpolator
func NewInterpolator() Interpolator {
	return interpolator{}
}

// SubstituteValues replaces all instances of 'key' with its respective 'value' from the `source` map in the `target` string and returns the applied target `string`.
func (i interpolator) SubstituteValues(source map[string]string, target string) string {
	signal := log.StarTrace("SubstituteValues")
	defer log.StopTrace(signal, nil)

	result := target
	hashedMap := map[string]string{}
	for key, value := range source {
		hasher := sha1.New()
		hasher.Write([]byte(key))
		hashedKey := hex.EncodeToString(hasher.Sum(nil))
		result = strings.ReplaceAll(result, fmt.Sprintf("%s%s", InterpolatorSignaler, key), fmt.Sprintf("%s{%s}", InterpolatorSignaler, hashedKey))
		hashedMap[hashedKey] = value
	}
	for key, value := range hashedMap {
		log.Debugf("Replaced %s[%s]", key, value)
		result = strings.ReplaceAll(result, fmt.Sprintf("%s{%s}", InterpolatorSignaler, key), value)
	}

	log.Debugf("Full replace string: %s", result)
	return result
}

// SubstituteValuesInMap replaces all instances of 'key' with its respective 'value' from the `source` map in the `target` map and returns the applied target `map`.
func (i interpolator) SubstituteValuesInMap(source map[string]string, target map[string]string) map[string]string {
	signal := log.StarTrace("SubstituteValuesInMap")
	defer log.StopTrace(signal, nil)

	interpolatedMap := map[string]string{}
	for targetKey, targetValue := range target {
		result := i.SubstituteValues(source, targetValue)

		interpolatedMap[targetKey] = result
	}

	return interpolatedMap
}
