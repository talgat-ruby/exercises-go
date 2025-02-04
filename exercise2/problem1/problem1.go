package problem1

func isChangeEnough(coins [4]int, amount float32) bool {
	total := int(amount * 100)
	nominal := []int{25, 10, 5, 1}
	for i, num := range coins {
		if num == 0 {
			continue
		}
		if needed := total / nominal[i]; num >= needed {
			total -= needed * nominal[i]
		} else {
			total -= num * nominal[i]
		}
	}

	return total == 0
}
