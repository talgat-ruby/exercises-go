package problem1

var quarter float32 = 0.25
var dime float32 = 0.1
var nickel float32 = 0.05
var penny float32 = 0.01

func isChangeEnough(changes [4]int, total float32) bool {

	return float32(changes[0])*quarter+float32(changes[1])*dime+float32(changes[2])*nickel+float32(changes[3])*penny >= total
}
