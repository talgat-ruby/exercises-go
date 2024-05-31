package problem11

import (
	"slices"
)

func keysAndValues[K int | string, V any](input map[K]V) ([]K, []V) {
	var keys []K
	var values []V

	for k := range input {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	for _, k := range keys {
		values = append(values, input[k])
	}

	return keys, values
}
