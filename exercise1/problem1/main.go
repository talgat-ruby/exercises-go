package main

func addUp(number int) int {
	result := 0
	for i := 1; i <= number; i++ {
		result = result + i
	}
	return result
}
