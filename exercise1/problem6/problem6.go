package problem6

func sumOfTwo(a []int, b []int, v int) bool {
	for _, av := range a {
		for _, bv := range b {
			if (av + bv) == v {
				return true
			}
		}
	}
	return false
}
