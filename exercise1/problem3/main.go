package main

func numberSquares(x int) int {
	var sum int = 0
	for i := range x{
		j := i + 1
		sum += (j * j)
	}
	return sum
}
