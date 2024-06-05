package problem11

func removeDups[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	var list []T
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return list
}
