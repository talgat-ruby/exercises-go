package problem11

import (
	"cmp"
	"sort"
)

func keysAndValues[K cmp.Ordered, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return cmp.Compare(keys[i], keys[j]) < 0
	})

	values := make([]V, 0, len(m))
	for _, k := range keys {
		values = append(values, m[k])
	}

	return keys, values
}
