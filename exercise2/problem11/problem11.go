package problem11

func removeDups[T comparable](items []T) []T {
	unique := make(map[T]bool)
	var result []T

	for _, item := range items {
		if !unique[item] {
			unique[item] = true
			result = append(result, item)
		}
	}
	return result
}
