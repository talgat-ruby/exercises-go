package problem1

func isChangeEnough(arr [4]int, coast float32) bool {
	sum := arr[0]*25 + arr[1]*10 + arr[2]*5 + arr[3]*1
	cena := coast * 100
	return sum >= int(cena)
}
