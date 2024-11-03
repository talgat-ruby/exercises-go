package problem11

func removeDups[T comparable](list []T) []T {
	seen := map[T]struct{}{}
	output := make([]T, 0, len(list))
	for _, v := range list {
		if _, ok := seen[v]; ok {
			continue
		}
		output = append(output, v)
		seen[v] = struct{}{}
	}

	return output
}
