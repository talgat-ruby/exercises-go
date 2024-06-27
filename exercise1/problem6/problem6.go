package problem6

func sumOfTwo(first []int, second []int, sum int) bool {
	for _, f := range first {
		for _, s := range second {
			if f+s == sum {
				return true
			}
		}
	}

	return false
}
