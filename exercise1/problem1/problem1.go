package problem1

func isChangeEnough(change [4]int, total float32) bool {
	totalCentsInteger := int(total * 100)
	if (change[0]*25 + change[1]*10 + change[2]*5 + change[3]) >= totalCentsInteger {
		return true
	}
	return false
}
