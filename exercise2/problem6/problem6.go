package problem6

func sumOfTwo(arr1 []int, arr2 []int, target int) bool {
	for _, value1 := range arr1 {
		for _, value2 := range arr2 {
			if (value1 + value2) == target {
				return true
			}
		}
	}
	return false
}
