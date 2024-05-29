package problem11

func removeDups[T comparable](input []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, v := range input {
		if _, exists := seen[v]; !exists {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
