package main

func addUp(num int) int {
	out := 0
	for i := 1; i <= num; i++ {
		out += i
	}
	return out
}
