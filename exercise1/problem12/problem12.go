package problem11

func keysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {

	var keys []K
	var values []V

	for key := range m {
		keys = append(keys, key)
		values = append(values, m[key])
	}

	return keys, values

}
