package problem11

func keysAndValues[K comparable, V any](input map[K]V) ([]K, []V) {
	var keys []K
	var values []V

	for k, v := range input {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}
