package problem11

func keysAndValues[K comparable, V any](input map[K]V) ([]K, []V) {
	keys := make([]K, len(input))
	values := make([]V, len(input))
	i := 0
	for k, v := range input {
		keys[i] = k
		values[i] = v
		i++
	}
	return keys, values
}
