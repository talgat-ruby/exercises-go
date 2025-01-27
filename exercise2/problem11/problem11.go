package problem11

func removeDups[T int | bool | string](items []T) []T {
	var result []T
	var check = make(map[T]bool)
	for _, item := range items {
		if !check[item] {
			result = append(result, item)
			check[item] = true
		}
	}

	return result
}
