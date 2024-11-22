package problem6

func sumOfTwo(a []int, b []int, target int) bool {

	bMap := make(map[int]struct{})

	for _, num := range b {
		bMap[num] = struct{}{}
	}

	for _, num := range a {
		complement := target - num
		if _, exists := bMap[complement]; exists {
			return true
		}
	}

	return false
}
