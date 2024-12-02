package problem6

func sumOfTwo(left, right []int, sum int) bool {
	mapper := make(map[int]struct{})
	for _, l := range left {
		mapper[sum-l] = struct{}{}
	}
	for _, r := range right {
		if _, ok := mapper[r]; ok {
			return true
		}
	}
	return false
}
