package problem6

func sumOfTwo(firstSlice []int, secondSlice []int, total int) bool {
	for _, number1 := range firstSlice {
		for _, number2 := range secondSlice {
			if  number1 + number2 == total{
				return true
			}
		}
	}

	return false
}
