package main

import "fmt"

func addUp(a int) int {
	sum := 0
	for c := 1; c <= a; c++ {
		sum = sum + c
	}
	return sum
}

func main() {
	var b int
	fmt.Scan(&b)
	res := addUp(b)
	fmt.Println(res)
}
