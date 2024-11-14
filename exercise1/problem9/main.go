package main

import "fmt"

func main() {
	var i int
	var j int
	i = 6
	j = 23
	fmt.Printf("Number %d and %d in binary is %b\n", i, j, i|j)
	fmt.Printf("Number %d and %d in binary is %b\n", i, j, i&j)
	fmt.Printf("Number %d and %d in binary is %b\n", i, j, i^j)
}
