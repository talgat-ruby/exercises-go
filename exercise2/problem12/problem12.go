package problem12

import (
	"fmt"
	"sort"
)

func keysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprint(keys[i]) < fmt.Sprint(keys[j])
	})

	for _, k := range keys {
		values = append(values, m[k])
	}

	return keys, values
}

func main() {
	keys, values := keysAndValues(map[string]int{"b": 2, "a": 1, "c": 3})
	fmt.Println(keys, values)

	keysStr, valuesStr := keysAndValues(map[string]string{"b": "Microsoft", "a": "Apple", "c": "Google"})
	fmt.Println(keysStr, valuesStr)

	keysInt, valuesBool := keysAndValues(map[int]bool{3: false, 1: true, 2: false})
	fmt.Println(keysInt, valuesBool)
}
