package problem11

import (
	"golang.org/x/exp/constraints"
	"sort"
)

func keysAndValues[K constraints.Ordered, V any](m map[K]V) ([]K, []V) {
	var keys []K
	var values []V
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	sortSlice(keys, values)

	return keys, values
}
func sortSlice[T constraints.Ordered, V any](k []T, v []V) {
	indices := make([]int, len(k))
	for i := range k {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return k[indices[i]] < k[indices[j]]
	})

	kCopy := make([]T, len(k))
	vCopy := make([]V, len(v))

	for i, idx := range indices {
		kCopy[i] = k[idx]
		vCopy[i] = v[idx]
	}

	copy(k, kCopy)
	copy(v, vCopy)
}
