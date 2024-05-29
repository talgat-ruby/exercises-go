package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	quarters := float32(changes[0])
	dimes := float32(changes[1])
	nickels := float32(changes[2])
	pennies := float32(changes[3])

	myTotal := quarters * 0.25 + dimes * 0.1 + nickels * 0.05 + pennies * 0.01

	if myTotal >= total {
		return true
	} else {
		return false
	}
}
