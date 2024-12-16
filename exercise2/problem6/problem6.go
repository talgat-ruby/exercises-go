package problem6

func sumOfTwo(a []int, b []int, s int) bool {
	for _, av := range a {
		for _, bv := range b {
			if av+bv == s {
				return true
			}
		}
	}
	return false
}
