package problem6

func sumOfTwo(a []int, b []int, c int) bool {
	for _, i := range a {
		for _, j := range b {
			if i+j == c {
				return true
			}
		}
	}
	return false
}
