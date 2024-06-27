package problem1

func isChangeEnough(change [4]int, total float32) bool {

	changeLeft := (float32(change[0]) * 0.25) + (float32(change[1]) * 0.1) + (float32(change[2]) * 0.05) + (float32(change[3]) * 0.01)

	diff := changeLeft - total
	if diff < 0.0 {
		return false
	}
	return true
}
