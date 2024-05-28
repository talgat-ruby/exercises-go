package problem6

func sumOfTwo(a, b []int, equal int) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i]+b[j] == equal {
				return true
			}
		}
	}
	return false
}
