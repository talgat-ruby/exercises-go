package problem6

func sumOfTwo(a []int, b []int, v int) bool {
	set := make(map[int]bool)

	for _, num := range a {
		set[v-num] = true
	}

	for _, num := range b {
		if set[num] {
			return true
		}
	}

	return false
}
