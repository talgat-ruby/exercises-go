package main

func addUp(number int) int {
	var total int
	for i := int(1); i <= number; i++ {
		total += i
	}
	return total
}
