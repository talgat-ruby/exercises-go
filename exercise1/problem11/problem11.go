package problem11

func removeDups[T comparable](list []T) []T {
	unique := map[T]struct{}{}
	var result []T
	for _, v := range list {
		if _, ok := unique[v]; !ok {
			result = append(result, v)
		}
		unique[v] = struct{}{}
	}
	return result
}
