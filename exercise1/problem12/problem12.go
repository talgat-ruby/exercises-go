package problem11

func keysAndValues[K comparable, V any](inp map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(inp))
	values := make([]V, 0, len(inp))
	for k, v := range inp {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
