package problem6

func sumOfTwo(a []int, b []int, target int) bool {
	for _, aV := range a {
		for _, bV := range b {
			if aV+bV == target {
				return true
			}
		}
	}
	return false
}
