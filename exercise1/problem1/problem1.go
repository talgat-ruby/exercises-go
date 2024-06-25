package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	notes := [4]float32{0.25, 0.10, 0.05, 0.01}
	var newArray [4]float32

	for i := 0; i < 4; i++ {
		newArray[i] = float32(changes[i]) * notes[i]
	}

	var sum float32
	for _, value := range newArray {
		sum += value
	}

	return sum >= total

}
