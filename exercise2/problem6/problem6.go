package problem6

func sumOfTwo(a []int, b []int, sum int) bool {
	result := false

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i]+b[j] == sum {
				result = true
			}
		}
	}

	return result
}
