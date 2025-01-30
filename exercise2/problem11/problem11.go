package problem11

func removeDups[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	var result []T

	for _, v := range input {
		if _, found := seen[v]; !found {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
