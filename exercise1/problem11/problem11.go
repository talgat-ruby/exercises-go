package problem11

func removeDups[T comparable](oldSlice []T) []T {
	var newSlice []T
	for i := 0; i < len(oldSlice); i++ {
		if !contains[T](newSlice, oldSlice[i]) {
			newSlice = append(newSlice, oldSlice[i])
		}
	}
	return newSlice
}

func contains[T comparable](list []T, item T) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == item {
			return true
		}
	}
	return false
}
