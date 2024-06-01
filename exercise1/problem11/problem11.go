package problem11

func removeDups[T comparable](input []T) []T {
	unique := make(map[T]struct{})
	result := []T{}
	for _, item := range input {
		if _, exists := unique[item]; !exists {
			unique[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
