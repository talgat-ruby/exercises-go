package problem11

func removeDups[T comparable](slice []T) []T {
	filter := make(map[T]bool)
	out := make([]T, 0, len(slice))
	for _, v := range slice {
		if _, found := filter[v]; !found {
			out = append(out, v)
			filter[v] = false
		}
	}
	return out
}
