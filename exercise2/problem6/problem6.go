package problem6

func sumOfTwo(a []int, b []int, sum int) bool {
	for _, val1 := range a {
		for _, val2 := range b {
			if val1+val2 == sum {
				return true
			}
		}
	}
	return false
}
