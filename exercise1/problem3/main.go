package main

import "fmt"

func numberSquares(st int) int {
	if st == 1 {
		return 1
	}
	b := st*st + numberSquares(st-1)
	return b
}

func main() {
	var b int
	fmt.Scan(&b)
	res := numberSquares(b)
	fmt.Print(res)
}
