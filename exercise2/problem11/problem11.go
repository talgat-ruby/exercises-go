package problem11

func removeDups[T comparable](items []T) []T {
	original := make(map[T]struct{})
	result := []T{}
	for _, item := range items {
		if _, found := original[item]; !found {
			original[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
