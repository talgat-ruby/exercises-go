package problem11

func removeDups[T comparable](input []T) []T {
	unique := make(map[T]bool)
	result := make([]T, 0)

	for _, item := range input {
		if !unique[item] {
			unique[item] = true
			result = append(result, item)
		}
	}

	return result
}
