package main

func numberSquares(num int) int {
	totalSquares := 0
	for k := 1; k <= num; k++ {
		totalSquares += (num - k + 1) * (num - k + 1)
	}
	return totalSquares
}
