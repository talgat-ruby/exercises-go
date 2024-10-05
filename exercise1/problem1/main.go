package main

func addUp(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		count += i
	}
	return count
}
