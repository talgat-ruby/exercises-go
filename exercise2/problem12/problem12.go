package problem11

import (
	"fmt"
	"sort"
)

func keysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprint(keys[i]) < fmt.Sprint(keys[j])
	})

	values := make([]V, 0, len(m))
	for _, key := range keys {
		values = append(values, m[key])
	}

	return keys, values
}
