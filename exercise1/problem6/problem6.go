package problem6

func sumOfTwo(a []int, b []int, target int) bool {
	for _, i := range a {
		for _, j := range b {
			if i+j == target {
				return true
			}
		}
	}
	return false
}
