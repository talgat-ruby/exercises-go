package main

import (
	"fmt"
)

func binary(de int) int {
	a := uint64(de)
	return int(a)
}

func main() {
	var c int
	fmt.Scan(&c)
	fmt.Printf("%b", c)
}
