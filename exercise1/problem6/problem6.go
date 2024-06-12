package problem6

func sumOfTwo(num1 []int, num2 []int, target int) bool {
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			if num1[i]+num2[j] == target {
				return true
			}
		}
	}
	return false
}
