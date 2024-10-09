package problem11

func removeDups[T comparable](slice []T) []T {
	encountered := make(map[T]bool)
	result := []T{}
	for _, item := range slice {
		if !encountered[item] {
			encountered[item] = true
			result = append(result, item)
		}
	}
	return result
}
