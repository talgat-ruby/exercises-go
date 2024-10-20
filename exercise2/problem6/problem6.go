package problem6

func sumOfTwo(a []int, b []int, v int) bool {
	sumSet := make(map[int]struct{})

	for _, numA := range a {
		sumSet[v-numA] = struct{}{}
	}

	for _, numB := range b {
		if _, exists := sumSet[numB]; exists {
			return true
		}
	}

	return false
}
