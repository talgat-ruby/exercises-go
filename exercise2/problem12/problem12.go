package problem11

func keysAndValues[K comparable, V comparable](inp map[K]V) ([]K, []V) {
	keys := []K{}
	values := []V{}
	for key, value := range inp {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}
