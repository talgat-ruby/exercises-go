package problem11

func removeDups[T comparable](inp []T) []T {
	seen := make(map[T]struct{})
	var result []T

	for _, v := range inp {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
