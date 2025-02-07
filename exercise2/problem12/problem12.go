package problem11

import (
	"fmt"
	"sort"
)

func keysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j])
	})
	sortedValues := make([]V, len(keys))
	for i, k := range keys {
		sortedValues[i] = m[k]
	}

	return keys, sortedValues
}
