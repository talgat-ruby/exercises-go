package problem11

import "sort"

func keysAndValues[K comparable, V any](input map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(input))
	values := make([]V, len(input))

	for k := range input {
		keys = append(keys, k)
	}

	asd := sort.MapKeys(input)

	for i, k := range keys {
		values[i] = input[k]
	}

	return keys, values
}
