package problem1

func isChangeEnough(cash [4]int, amount float32) bool {

	allCash := cash[0]*25 + cash[1]*10 + cash[2]*5 + cash[3]
	amountToCent := int(amount * 100)

	return allCash >= amountToCent
}
