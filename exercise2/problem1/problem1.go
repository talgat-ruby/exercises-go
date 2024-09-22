package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	cent := []int{25, 10, 5, 1}

	sumCent := int(total * 100)
	result := 0
	for i, count := range changes {
		result += count * cent[i]
	}
	return result >= sumCent
}
