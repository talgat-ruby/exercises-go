package problem11

import (
	"fmt"
	"sort"
)

func keysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	if m == nil {
		return []K{}, []V{}
	}

	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return stringify(keys[i]) < stringify(keys[j])
	})

	values := make([]V, len(keys))
	for i, k := range keys {
		values[i] = m[k]
	}

	return keys, values
}

func stringify[T any](v T) string {
	return fmt.Sprintf("%v", v)
}
