package main

func addUp(num int) int {
	var result int
	for i := 1; i <= num; i++ {
		result = result + i
	}
	return result
}
