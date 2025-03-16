package problem11

import (
	"slices"
)

func keysAndValues[T int | string, P int | bool | string](m map[T]P) ([]T, []P) {

	var keys []T
	var values []P
	for key, value := range m {
		keys = append(keys, key)
		values = append(values, value)
	}

	slices.SortFunc(keys, func(a, b T) int {
		if a > b {
			return 1
		}
		return -1
	})
	return keys, values
}
