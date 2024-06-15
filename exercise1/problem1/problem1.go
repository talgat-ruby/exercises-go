package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	quarters := float32(changes[0]) * 0.25
	dime := float32(changes[1]) * 0.10
	nickel := float32(changes[2]) * 0.05
	pennies := float32(changes[3]) * 0.01

	if quarters+dime+nickel+pennies >= total {
		return true
	}
	return false
}
