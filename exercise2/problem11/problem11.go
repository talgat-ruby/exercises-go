package problem11

func removeDups[T any](array []T) []T {
	out := make([]T, 0, len(array))
	filter := make(map[any]struct{})
	for _, val := range array {
		if _, ok := filter[val]; !ok {
			filter[val] = struct{}{}
			out = append(out, val)
		}
	}
	return out
}
