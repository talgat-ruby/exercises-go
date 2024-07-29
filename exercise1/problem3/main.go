package main

func numberSquares(size int) int {
	count := 0
	for i := 1; i <= size; i++ {
		count += i * i
	}
	return count
}
