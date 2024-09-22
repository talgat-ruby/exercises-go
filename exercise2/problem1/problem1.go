package problem1

func isChangeEnough(amount float64, coins []int) bool {
	cents := []int{25, 10, 5, 1}

	amountCents := int(amount * 100)

	totalCent := 0

	for i, count := range coins {
		totalCent += count * cents[i]
	}
	return totalCent >= amountCents
}
