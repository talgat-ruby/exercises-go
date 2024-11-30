package problem1

func isChangeEnough(pocket [4]int, item float32) bool {
	if item < 0 {
		return false
	}

	for _, coins := range pocket {
		if coins < 0 {
			return false
		}
	}

	diff := item - (0.25*float32(pocket[0]) + 0.1*float32(pocket[1]) + 0.05*float32(pocket[2]) + 0.01*float32(pocket[3]))
	if diff <= 0 {
		return true
	}
	return false
}
