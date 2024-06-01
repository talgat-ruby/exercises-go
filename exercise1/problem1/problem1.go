package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	var pocket = [4]float32{0.25, 0.10, 0.05, 0.01}
	var sum = float32(0.0)

	for i := 0; i < len(changes); i++ {
		sum += float32(changes[i]) * pocket[i]
	}

	return sum >= total
}
