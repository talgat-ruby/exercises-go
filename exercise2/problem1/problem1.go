package problem1

func isChangeEnough(change [4]int, total float32) bool {
	quarters := float32(change[0]) * 0.25
	dimes := float32(change[1]) * 0.10
	nickels := float32(change[2]) * 0.05
	pennies := float32(change[3]) * 0.01

	totalChange := quarters + dimes + nickels + pennies

	return totalChange >= total
}
