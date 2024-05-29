package problem11

func keysAndValues[T ~int | ~string | ~bool, P ~int | ~string | ~bool](itemsMap map[T]P) ([]T, []P) {
	keys := []T{}
	vals := []P{}
	for k, v := range itemsMap {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	return keys, vals
}
