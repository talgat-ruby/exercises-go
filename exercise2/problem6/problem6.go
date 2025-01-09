package problem6

func sumOfTwo(arr1 []int, arr2 []int, target int) bool {
	for _, el1 := range arr2 {
		for _, el2 := range arr1 {
			if el1+el2 == target {
				return true
			}
		}
	}
	return false
}
