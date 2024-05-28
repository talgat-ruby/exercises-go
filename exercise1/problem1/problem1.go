package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	var sum float32
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			sum += float32(changes[i]) * 0.25
		case 1:
			sum += float32(changes[i]) * 0.10
		case 2:
			sum += float32(changes[i]) * 0.05
		case 3:
			sum += float32(changes[i]) * 0.01
		}
	}
	return sum >= total
}
