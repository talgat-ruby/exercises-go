package problem6

func sumOfTwo(a []int, b []int, v int) bool {

	complementSet := make(map[int]bool)

	for _, num := range a {
		complementSet[v-num] = true
	}

	for _, num := range b {
		if complementSet[num] {
			return true
		}
	}

	return false
}
