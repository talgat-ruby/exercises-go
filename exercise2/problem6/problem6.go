package problem6

func sumOfTwo(a []int, b []int, target int) bool {
	complements := make(map[int]bool)
	for _, num := range a {
		complements[target-num] = true
	}
	for _, num := range b {
		if complements[num] {
			return true
		}
	}

	return false
}
