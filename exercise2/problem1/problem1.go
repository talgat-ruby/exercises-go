package problem1

func isChangeEnough(coins [4]int, cost float32) bool {
	return (float32(coins[0])*0.25 + float32(coins[1])*0.10 + float32(coins[2])*0.05 + float32(coins[3])*0.01) >= cost
}
