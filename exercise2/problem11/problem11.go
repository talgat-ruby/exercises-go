package problem11

func removeDups[T comparable](items []T) []T {
	output := []T{}
	tempMap := map[T]bool{}

	for _, item := range items {
		if !tempMap[item] {
			output = append(output, item)
			tempMap[item] = true
		}
	}

	return output
}
