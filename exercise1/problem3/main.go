package main

func numberSquares(x int) int {
	var ac int
	for x > 0 {
		ac += x * x
		x -= 1
	}
	return ac
}
