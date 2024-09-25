package problem11

func removeDups[T comparable](input []T) []T {
	var result []T
	dups := make(map[T]bool)
	for _, v := range input {
		if !dups[v] {
			dups[v] = true
			result = append(result, v)
		}
	}
	return result
}
