package problem1

import (
	"fmt"
)

func isChangeEnough(change [4]int, amountDue float32) bool {

	totalDueCents := int(amountDue * 100)

	totalChangeCents := change[0]*25 + change[1]*10 + change[2]*5 + change[3]*1

	return totalChangeCents >= totalDueCents
}

func main() {
	fmt.Println(isChangeEnough([4]int{2, 100, 0, 0}, 14.11))
	fmt.Println(isChangeEnough([4]int{0, 0, 20, 5}, 0.75))
	fmt.Println(isChangeEnough([4]int{30, 40, 20, 5}, 12.55))
	fmt.Println(isChangeEnough([4]int{10, 0, 0, 50}, 3.85))
	fmt.Println(isChangeEnough([4]int{1, 0, 5, 219}, 19.99))
}
