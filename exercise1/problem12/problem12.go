package problem11

import (
	"fmt"
	"sort"
)

func keysAndValues[K comparable, V any](input map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(input))
	values := make([]V, 0, len(input))

	for key := range input {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprint(keys[i]) < fmt.Sprint(keys[j])
	})

	for _, key := range keys {
		values = append(values, input[key])
	}

	return keys, values
}
