package problem11

import (
	"sort"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func keysAndValues[K Ordered, V any](m map[K]V) ([]K, []V) {
	var keys []K
	var values []V
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	sortSlice(keys, values)

	return keys, values
}
func sortSlice[T Ordered, V any](k []T, v []V) {
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
