package problem6

func sumOfTwo(f []int, s []int, sum int) bool {
	for _, v := range f {
		for _, w := range s {
			if v+w == sum {
				return true
			}
		}
	}
	return false
}
