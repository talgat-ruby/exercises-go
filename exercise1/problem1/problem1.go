package problem1

func isChangeEnough(changes [4]int, totalDue float32) bool {
	totalChange := float32(changes[0]*0.25 + changes[1]*0.10 + changes[2]*0.05 + changes[3]*0.01)
	return totalChange >= totalDue
}
