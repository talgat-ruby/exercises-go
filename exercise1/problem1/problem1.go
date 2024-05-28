package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	quarters := float32(changes[0])
	dimes := float32(changes[1])
	nickles := float32(changes[2])
	pennies := float32(changes[3])

	return (quarters*0.25 + dimes*0.10 + nickles*0.05 + pennies*0.01) >= total
}
