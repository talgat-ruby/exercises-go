package main

func addUp(num int) int {
	result := 0
	for i := 1; i <= num; i++ {
		result += i
	}
	return result
}
