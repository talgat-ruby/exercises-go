package main

func numberSquares(num int) int {
	if num == 1 {
		return num
	}
	return num*num + numberSquares(num-1)
}
