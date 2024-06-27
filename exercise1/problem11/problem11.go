package problem11

func removeDups[V comparable](values []V) []V {
	uniqueMap := make(map[V]bool)
	res := make([]V, 0)
	for _, v := range values {
		if !uniqueMap[v] {
			res = append(res, v)
		}
		uniqueMap[v] = true
	}

	return res
}
