package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	coinValues := [4]float32{0.25, 0.10, 0.05, 0.01}
	var dollars float32

	for i, coin := range changes {
		dollars += float32(coin) * coinValues[i]
	}

	return dollars >= total
}
