package problem1

func isChangeEnough(changes [4]int, total float32) bool {

	var quarter, dime, nickel, penny float32
	ischange := changes
	var sum float32

	quarter = 0.25
	dime = 0.1
	nickel = 0.05
	penny = 0.01
	sum = float32(ischange[0]) * quarter
	sum = sum + float32(ischange[1])*dime
	sum = sum + float32(ischange[2])*nickel
	sum = sum + float32(ischange[3])*penny

	if total <= sum {
		return true
	} else {
		return false
	}

}
