package main

func numberSquares(num int) int {
	numOfSquares := 0
	for x := range num + 1 {
		numOfSquares = numOfSquares + x*x
	}
	return numOfSquares
}
