package problem6

func sumOfTwo(first []int, second []int, target int) bool {
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			if first[i]+second[j] == target {
				return true
			}
		}
	}
	return false
}
