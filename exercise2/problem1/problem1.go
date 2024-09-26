package problem1

func isChangeEnough(arr [4]int, price float32) bool {

	sum := (0.25)*float32(arr[0]) + (0.1)*float32(arr[1]) + (0.05)*float32(arr[2]) + (0.01)*float32(arr[3])
	if sum >= price {
		return true
	}
	return false
}
