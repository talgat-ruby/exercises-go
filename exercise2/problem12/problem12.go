package problem11

import "fmt"

func keysAndValues[K comparable, V any](map_keysAndValues map[K]V) ([]K, []V) {
	mas_keys := make([]K, 0, len(map_keysAndValues))
	mas_values := make([]V, 0, len(map_keysAndValues))

	for k, v := range map_keysAndValues {
		mas_keys = append(mas_keys, k)
		mas_values = append(mas_values, v)
	}

	sort_keys(mas_keys)

	sortedValues := make([]V, len(mas_keys))
	for i, k := range mas_keys {
		sortedValues[i] = map_keysAndValues[k]
	}

	return mas_keys, sortedValues
}

func sort_keys[K comparable](keys []K) {
	n := len(keys)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if fmt.Sprintf("%v", keys[j]) > fmt.Sprintf("%v", keys[j+1]) {
				keys[j], keys[j+1] = keys[j+1], keys[j]
			}
		}
	}
}
