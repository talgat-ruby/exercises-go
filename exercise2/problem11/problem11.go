package problem11

func removeDups[T comparable](items []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, item := range items {
		if !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}

	return result
}
