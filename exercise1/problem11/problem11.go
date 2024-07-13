package problem11

func removeDups[T ~int | ~bool | ~string](items []T) []T {
	itemsMap := make(map[T]int)
	uniqueItems := []T{}
	for _, item := range items {
		_, found := itemsMap[item]
		if !found {
			uniqueItems = append(uniqueItems, item)
			itemsMap[item] += 1
		}
	}
	return uniqueItems
}
