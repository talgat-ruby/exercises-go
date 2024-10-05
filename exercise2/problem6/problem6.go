package problem6

func sumOfTwo(a, b []int, sum int) bool {
	for _, i := range a {
		for _, j := range b {
			if i+j == sum {
				return true
			}
		}
	}
	return false
}
