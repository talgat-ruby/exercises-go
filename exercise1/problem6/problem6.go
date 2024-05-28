package problem6

func sumOfTwo(sliceA, sliceB []int, sum int) bool {
	mapA := make(map[int]int)
	for i, v := range sliceA {
		if sum > v {
			mapA[sum-v] = i
			for _, e := range sliceB {
				if _, ok := mapA[e]; ok {
					return true
				}
			}
		}
	}
	return false
}
