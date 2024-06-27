package problem1

func isChangeEnough(coins [4]int, total float32) bool {
	values := []float32{0.25, 0.10, 0.05, 0.01}
	for i, nominal := range values {
		if total < nominal {
			continue
		}
		requiredCoins := int(total / nominal)
		if requiredCoins > coins[i] && i == 3 {
			return false
		} else if requiredCoins > coins[i] {
			requiredCoins = coins[i]
		}
		total -= float32(requiredCoins) * nominal
	}
	return int(total) == 0
}
