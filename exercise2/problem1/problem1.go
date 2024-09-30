package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	value := float32(changes[0])*0.25 + float32(changes[1])*0.10 + float32(changes[2])*0.05 + float32(changes[3])*0.01

	return value >= total
}
