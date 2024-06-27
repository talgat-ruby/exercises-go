package problem12

import (
	"cmp"
	"sort"
)

func keysAndValues[K cmp.Ordered, V any](inp map[K]V) ([]K, []V) {
	var keys []K
	for k, _ := range inp {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var values []V
	for _, k := range keys {
		values = append(values, inp[k])
	}

	return keys, values
}
