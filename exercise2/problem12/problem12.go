package problem11

import "slices"

func keysAndValues[K ~int | ~string, V ~int | ~string | ~bool](items map[K]V) ([]K, []V) {
	if len(items) == 0 {
		return []K{}, []V{}
	}
	keys := make([]K, 0, len(items))
	values := make([]V, 0, len(items))

	for key := range items {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	//sort.Slice(keys, func(i, j int) bool {
	//	return keys[i] < keys[j]
	//})

	for _, key := range keys {
		values = append(values, items[key])
	}
	return keys, values

}
