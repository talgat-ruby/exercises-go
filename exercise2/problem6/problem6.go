package problem6

func sumOfTwo(first []int, second []int, sum int) bool {
	for _, fv := range first {
		for _, sv := range second {
			if fv+sv == sum {
				return true
			}
		}
	}
	return false
}
