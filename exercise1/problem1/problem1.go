package problem1

func isChangeEnough(change [4]int, amount float32) bool {
	total := float32(change[0])*0.25 + float32(change[1])*0.10 + float32(change[2])*0.05 + float32(change[3])*0.01
	return total >= amount
}
