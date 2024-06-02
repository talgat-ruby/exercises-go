package problem11

import (
	"cmp"
	"slices"
)

func keysAndValues[K cmp.Ordered, V any](inMap map[K]V) (keys []K, values []V) {
	for k := range inMap {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		values = append(values, inMap[k])
	}
	return keys, values
}
