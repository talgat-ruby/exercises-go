package problem11

import (
	"fmt"
	"sort"
)

func keysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j])
	})

	values := make([]V, len(keys))
	for i, k := range keys {
		values[i] = m[k]
	}

	return keys, values
}

func main() {

	keys1, values1 := keysAndValues(map[string]int{"a": 1, "b": 2, "c": 3})
	fmt.Println(keys1, values1)

	keys2, values2 := keysAndValues(map[string]string{"a": "Apple", "b": "Microsoft", "c": "Google"})
	fmt.Println(keys2, values2)

	keys3, values3 := keysAndValues(map[int]bool{1: true, 2: false, 3: false})
	fmt.Println(keys3, values3)
}
