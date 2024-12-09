package problem12

import (
	"cmp"
	"slices"
)

func keysAndValues[K cmp.Ordered, V any](m map[K]V) ([]K, []V) {
	if m == nil {
		return []K{}, []V{}
	}

	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	values := make([]V, len(keys))
	for i, k := range keys {
		values[i] = m[k]
	}

	return keys, values
}
