package io

// MergeMap joins source map with target map if the key exists on both source overrides the value from target
func MergeMap(source map[string]interface{}, target map[string]interface{}) map[string]interface{} {
	for key, value := range source {
		target[key] = value
	}

	return target
}

// MergeStringMap joins source map with target map if the key exists on both source overrides the value from target
func MergeStringMap(source map[string]string, target map[string]string) map[string]string {
	for key, value := range source {
		target[key] = value
	}

	return target
}
