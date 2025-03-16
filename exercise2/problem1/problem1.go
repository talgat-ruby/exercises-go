package problem1

func isChangeEnough(s [4]int, price float32) bool {
	sum := float32(s[0])*0.25 + float32(s[1])*0.1 + float32(s[2])*0.05 + float32(s[3])*0.01
	return sum >= price
}
