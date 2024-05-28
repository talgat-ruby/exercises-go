package problem11

import (
	"fmt"
	"reflect"
	"sort"
)

func keysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	// Extract keys from the map.
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Sort the keys using reflection.
	sort.Slice(keys, func(i, j int) bool {
		vi := reflect.ValueOf(keys[i])
		vj := reflect.ValueOf(keys[j])
		switch vi.Kind() {
		case reflect.String:
			return vi.String() < vj.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return vi.Int() < vj.Int()
		default:
			// Add more cases if needed
			return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j])
		}
	})

	// Extract values based on the sorted keys.
	values := make([]V, 0, len(m))
	for _, k := range keys {
		values = append(values, m[k])
	}

	return keys, values
}
