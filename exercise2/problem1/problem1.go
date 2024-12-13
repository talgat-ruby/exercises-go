package problem1

var cents = []float32{0.25, 0.10, 0.05, 0.01}

func isChangeEnough(changes [4]int, total float32) bool {
	var sum float32
	for key, index := range changes {
		sum += cents[key] * float32(index)
	}
	return sum >= total
}
