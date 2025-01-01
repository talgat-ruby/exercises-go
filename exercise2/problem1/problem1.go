package problem1

func isChangeEnough(changes [4]int, total float32) bool {

	quarter := float32(changes[0]) * 0.25
	dime := float32(changes[1]) * 0.10
	nickel := float32(changes[2]) * 0.05
	penny := float32(changes[3]) * 0.01

	ChangeEnough := quarter + dime + nickel + penny

	return ChangeEnough >= total
}
