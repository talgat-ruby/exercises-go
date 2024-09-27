package main

import (
	"fmt"
)

func addUp(n int) int {
	return n * (n + 1) / 2
}

func main() {

	fmt.Println(addUp(4))
	fmt.Println(addUp(13))
	fmt.Println(addUp(600))
}
