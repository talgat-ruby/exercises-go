package problem11

func keysAndValues[K, V comparable](input map[K]V) ([]K, []V) {
	var keys []K
	var values []V
	for key, value := range input {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}
