package main

func addUp(num int) int {
	var count int
	for i := 1; i <= num; i++ {
		count += i
	}
	return count
}
