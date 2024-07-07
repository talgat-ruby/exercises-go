package problem6

func sumOfTwo(a []int, b []int, v int) bool {
	for _, v1 := range a {
		for _, v2 := range b {
			if v1+v2 == v {
				return true
			}
		}
	}

	return false
}
