package problem11

import (
	"golang.org/x/exp/constraints"
	"sort"
)

func keysAndValues[K constraints.Ordered, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0)
	values := make([]V, 0)

	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, v := range keys {
		values = append(values, m[v])
	}

	return keys, values
}
