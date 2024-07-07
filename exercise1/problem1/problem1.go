package problem1

func isChangeEnough(changes [4]int, amountToPay float32) bool {
	total := float32(changes[0]) * 0.25
	total += float32(changes[1]) * 0.10
	total += float32(changes[2]) * 0.05
	total += float32(changes[3]) * 0.01

	return float32(total) >= amountToPay
}
