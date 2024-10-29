package problem11

import (
	"fmt"
	"sort"
)

func keysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprint(keys[i]) < fmt.Sprint(keys[j])
	})

	values := make([]V, 0, len(m))
	for _, k := range keys {
		values = append(values, m[k])
	}

	return keys, values
}
