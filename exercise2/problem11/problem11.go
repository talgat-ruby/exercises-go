package problem11

func removeDups[T comparable](items []T) []T {

	seen := make(map[T]struct{})
	result := []T{}

	for _, item := range items {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}
