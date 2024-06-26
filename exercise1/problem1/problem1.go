package problem1

const (
	quarter = 0.25
	dime    = 0.1
	nickel  = 0.05
	penny   = 0.01
)

func isChangeEnough(changes [4]int, total float32) bool {
	totalChange := float32(changes[0])*quarter + float32(changes[1])*dime + float32(changes[2])*nickel + float32(changes[3])*penny
	result := totalChange >= total

	return result
}
