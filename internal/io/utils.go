package io

func mergeMap(source map[string]string, target map[string]string) map[string]string {
	for key, value := range source {
		target[key] = value
	}

	return target
}
