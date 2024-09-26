package problem6

func sumOfTwo(a []int, b []int, target int) bool {
	valueMap := make(map[int]bool)
	for _, value := range b {
		valueMap[value] = true
	}
	for _, value := range a {
		if valueMap[target-value] {
			return true
		}
	}
	return false

}
