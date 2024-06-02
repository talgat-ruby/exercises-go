package problem1

func isChangeEnough(change []int, amountDue float64) bool {
	quarterValue := 0.25
	dimeValue := 0.10
	nickelValue := 0.05
	pennyValue := 0.01

	// Рассчитываем общую сумму денег
	total := float64(change[0])*quarterValue + float64(change[1])*dimeValue + float64(change[2])*nickelValue + float64(change[3])*pennyValue

	// Сравниваем с долгом
	return total >= amountDue
}
