package problem1

func isChangeEnough(change [4]int, total float32) bool {
	quarterValue := 0.25
	dimeValue := 0.10
	nickelValue := 0.05
	pennyValue := 0.01

	totalChange := float32(change[0])*float32(quarterValue) + float32(change[1])*float32(dimeValue) +
		float32(change[2])*float32(nickelValue) + float32(change[3])*float32(pennyValue)

	return totalChange >= total
}
