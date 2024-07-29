package main

func numberSquares(num int) int {
	var result int
	for i := 1; i <= num; i++ {
		result = result + (i * i)
	}
	return result
}
