package problem11

func removeDups[T comparable](slice []T) []T {
	list := make(map[T]struct{})
	result := make([]T, 0)

	for _, key := range slice {
		if _, exists := list[key]; !exists {
			list[key] = struct{}{}
			result = append(result, key)
		}
	}
	return result
}
