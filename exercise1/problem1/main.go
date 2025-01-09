package main

func addUp(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	total := 0
	for i := num; i > 0; i-- {
		total += i
	}
	return total
}
