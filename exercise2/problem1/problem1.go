package problem1

func isChangeEnough(change [4]int, amountDue float32) bool {
	quartersValue := 25
	dimesValue := 10
	nickelsValue := 5
	pennieValue := 1

	totalCents := (change[0] * quartersValue) +
		(change[1] * dimesValue) +
		(change[2] * nickelsValue) +
		(change[3] * pennieValue)

	amountDueCents := int(amountDue * 100)

	return totalCents >= amountDueCents
}
