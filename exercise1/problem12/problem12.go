package problem11

import "slices"

func keysAndValues[K int | string, V any](m map[K]V) ([]K, []V) {

	var keys []K
	var values []V

	for key := range m {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	for _, key := range keys {
		values = append(values, m[key])
	}

	return keys, values

}
