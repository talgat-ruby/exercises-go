package problem11

import (
	"sort"
	"fmt"
)

func keysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Сортируем ключи
	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprint(keys[i]) < fmt.Sprint(keys[j])
	})

	values := make([]V, len(m))
	for i, k := range keys {
		values[i] = m[k]
	}

	return keys, values
}