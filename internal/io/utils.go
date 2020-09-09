package io

func mergeMap(left map[string]string, right map[string]string) map[string]string {
	for key, value := range right {
		left[key] = value
	}

	return left
}
