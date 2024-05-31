package problem11

func removeDups[T int | bool | string](input []T) []T {
	var result []T
	uniqueArr := make(map[T]bool)

	for _, item := range input {
		if !uniqueArr[item] {
			uniqueArr[item] = true
			result = append(result, item)
		}
	}

	return result
}
