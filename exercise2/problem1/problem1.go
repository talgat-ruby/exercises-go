package problem1

import (
	"fmt"
)

func isChangeEnough(coins [4]int, due float32) bool {
	quartersValue := 0.25
	dimesValue := 0.10
	nickelsValue := 0.05
	penniesValue := 0.01

	total := float64(coins[0])*quartersValue + float64(coins[1])*dimesValue +
		float64(coins[2])*nickelsValue + float64(coins[3])*penniesValue

	return total >= float64(due)
}

func main() {
	fmt.Println(isChangeEnough([4]int{2, 100, 0, 0}, 14.11))
	fmt.Println(isChangeEnough([4]int{0, 0, 20, 5}, 0.75))
	fmt.Println(isChangeEnough([4]int{30, 40, 20, 5}, 12.55))
	fmt.Println(isChangeEnough([4]int{10, 0, 0, 50}, 3.85))
	fmt.Println(isChangeEnough([4]int{1, 0, 5, 219}, 19.99))
}
