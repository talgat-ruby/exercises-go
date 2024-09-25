package problem6

func sumOfTwo(a []int, b []int, sum int) bool {
	for _, one := range a {
		for _, two := range b {
			if one+two == sum {
				return true
			}
		}
	}
	return false
}
