package problem6

func sumOfTwo(a []int, b []int, v int) bool {
	needed := make(map[int]struct{})

	for _, num := range b {
		needed[v-num] = struct{}{}
	}

	for _, num := range a {
		if _, found := needed[num]; found {
			return true
		}
	}

	return false
}
