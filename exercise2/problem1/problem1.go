package problem1

func isChangeEnough(changes [4]int, total float32) bool {

	var quarter, dime, nickel, penny float32
	ischange := changes
	var sum float32

	quarter = 0.25
	dime = 0.1
	nickel = 0.05
	penny = 0.01
	sum += float32(ischange[0]) * quarter
	sum += float32(ischange[1]) * dime
	sum += float32(ischange[2]) * nickel
	sum += float32(ischange[3]) * penny

	if total <= sum {
		return true
	} else {
		return false
	}

}

// quarter: 25 cents / $0.25
// dime: 10 cents / $0.10
// nickel: 5 cents / $0.05
// penny: 1 cent / $0.01
