package problem1

func IsChangeEnough(change []int, amount float64) bool {
	coinValues := []float64{0.25, 0.10, 0.05, 0.01}

	totalChange := 0.0
	for i := 0; i < len(change); i++ {
		totalChange += float64(change[i]) * coinValues[i]
	}
	return totalChange >= amount
}
