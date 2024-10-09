package problem11

import "sort"

func keysAndValues[K string | int, V comparable](mp map[K]V) ([]K, []V) {

	keys := make([]K, 0, len(mp))
	values := make([]V, 0, len(mp))

	for k := range mp {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		values = append(values, mp[k])
	}

	return keys, values
}
