package problem1

func isChangeEnough(pocket [4]int, coast float32) bool {
	var total float32
	total += float32(pocket[0]) * 0.25
	total += float32(pocket[1]) * 0.1
	total += float32(pocket[2]) * 0.05
	total += float32(pocket[3]) * 0.01
	return total >= coast
}
