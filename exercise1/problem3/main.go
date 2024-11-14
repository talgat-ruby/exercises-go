package main

import "fmt"

func main() {
	var a = numberSquares(2)
	var b = numberSquares(3)
	var c = numberSquares(4)
	var d = numberSquares(5)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func numberSquares(n int) int {
	return (n + 1) * n * (2*n + 1) / 6
}
