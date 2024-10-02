package problem11

func removeDups[T comparable](slice []T) []T {
	result := []T{}
	seen := map[T]struct{}{}
	for _, value := range slice {
		if _, ok := seen[value]; !ok {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}
	return result
}
