package problem11

func removeDups[T comparable](inp []T) []T {
	var result []T

	for i, t := range inp {
		if notContains(result, inp[i]) {
			result = append(result, t)
		}
	}

	return result
}

func notContains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return false
		}
	}
	return true
}
