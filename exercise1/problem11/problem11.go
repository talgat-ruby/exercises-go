package problem11

func removeDups[T comparable](input []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}

	for _, item := range input {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return list
}
