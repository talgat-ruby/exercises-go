package problem1

func isChangeEnough(values [4]int, totalValue float32) bool {
	quarters := float32(values[0]) * 0.25
	dimes := float32(values[1]) * 0.10
	nickels := float32(values[2]) * 0.05
	penny := float32(values[3]) * 0.01

	total := quarters + dimes + nickels + penny

	return total >= totalValue
}
