package problem6

func sumOfTwo(a []int, b []int, v int) bool {
	for _, numA := range a {
		for _, numB := range b {
			if numA+numB == v {
				return true
			}
		}
	}
	return false
}
