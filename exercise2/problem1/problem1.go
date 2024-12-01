package problem1

func isChangeEnough(arr [4]int, amount float32) bool {
	quarters := 25 * arr[0]
	dimes := 10 * arr[1]
	nickels := 5 * arr[2]
	pennies := 1 * arr[3]
	sum := (float32)(quarters + dimes + nickels + pennies)
	return sum-amount*100 > -1
}
