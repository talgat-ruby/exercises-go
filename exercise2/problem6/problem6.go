package problem6

func sumOfTwo(a []int, b []int, v int) bool {

	bMap := make(map[int]struct{})
	for _, num := range b {
		bMap[num] = struct{}{}
	}

	for _, num := range a {
		complement := v - num
		if _, exists := bMap[complement]; exists {
			return true
		}
	}

	return false
}
