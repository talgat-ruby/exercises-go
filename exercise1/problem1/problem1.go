package problem1

// HELLO
func isChangeEnough(change [4]int, totalDue float32) bool {
	sum := float32(change[0]) * 0.25
	sum += float32(change[1]) * 0.10
	sum += float32(change[2]) * 0.05
	sum += float32(change[3]) * 0.01

	if sum >= totalDue {
		return true
	} else {
		return false
	}
}
