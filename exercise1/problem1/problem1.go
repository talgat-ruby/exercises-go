package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	quarter := float32(changes[0])
	dime := float32(changes[1])
	nickel := float32(changes[2])
	penny := float32(changes[3])

	sum := quarter*0.25 + dime*0.10 + nickel*0.05 + penny*0.01
	return sum >= total
}
