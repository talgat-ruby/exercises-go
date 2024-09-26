package problem1

func isChangeEnough(change []int, total float64) bool {

	quarters := float64(change[0]) * 0.25
	dimes := float64(change[1]) * 0.10
	nickels := float64(change[2]) * 0.05
	pennies := float64(change[3]) * 0.01

	totalChange := quarters + dimes + nickels + pennies

	return totalChange >= total
}
