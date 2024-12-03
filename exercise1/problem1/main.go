package main

import (
	"fmt"
)

func addUp(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

func main() {
	fmt.Println(addUp(4))
	fmt.Println(addUp(13))
	fmt.Println(addUp(600))
}
