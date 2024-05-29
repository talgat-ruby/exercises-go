package problem11

func removeDups[T ~int | ~bool | ~string](items []T) []T {
	itemsMap := make(map[T]int)
	for _, item := range items {
		itemsMap[item] = 0
	}
	uniqueItems := []T{}
	for k, _ := range itemsMap {
		uniqueItems = append(uniqueItems, k)
	}
	return uniqueItems
}
