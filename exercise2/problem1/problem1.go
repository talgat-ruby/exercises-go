package problem1

func isChangeEnough(change [4]int, amount float32) bool {
	var (
		is_pay   bool = false
		summa    int  = 0
		nominals      = [4]int{25, 10, 5, 1}
	)
	for i, v := range change {
		summa += nominals[i] * v
	}
	amount_int := int(amount * 100)
	if summa >= amount_int {
		is_pay = true
	}
	return is_pay
}
