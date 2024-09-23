package main

func numberSquares(num int) int {
	out := 0
	for i := 1; i <= num; i++ {
		out += i * i
	}
	return out
}
