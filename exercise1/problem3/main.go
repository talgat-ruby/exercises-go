package main

func numberSquares(num int) int {

	squares := 0

	for i := 1; i <= num; i++ {
		squares += (i * i)
	}

	return squares
}
