package problem11

func removeDups[T comparable](input []T) []T {
	encountered := map[T]struct{}{}
	result := []T{}

	for _, value := range input {
		if _, ok := encountered[value]; !ok {
			encountered[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}
