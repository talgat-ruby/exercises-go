package problem6

func sumOfTwo(slc1 []int, slc2 []int, number int) bool {
	var is_sum bool
	for _, v1 := range slc1 {
		for _, v2 := range slc2 {
			if (v1 + v2) == number {
				is_sum = true
				break
			}
		}
	}
	return is_sum
}
