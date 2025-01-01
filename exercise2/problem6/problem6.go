package problem6

func sumOfTwo(a []int, b []int, sum int) bool {
	complements := make(map[int]bool)
	for _, numA := range a {
		complements[sum-numA] = true
	}
	for _, numB := range b {
		if complements[numB] {
			return true
		}
	}
	return false
}
