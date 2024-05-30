package problem11

func removeDups[T comparable](slice []T) []T {
	res := make([]T, 0, len(slice))
	for _, v := range slice {
		if contains(res, v) {
			continue
		}

		res = append(res, v)
	}

	return res
}

func contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if target == v {
			return true
		}
	}
	return false
}
