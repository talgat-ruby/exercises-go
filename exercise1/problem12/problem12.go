package problem11

import "sort"

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func keysAndValues[K Ordered, V any](inp map[K]V) ([]K, []V) {
	var keys []K
	var values []V
	for k, _ := range inp {
		keys = append(keys, k)
	}
	sortSlice(keys)
	for i := 0; i < len(keys); i++ {
		values = append(values, inp[keys[i]])
	}
	return keys, values
}

func sortSlice[T Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}
