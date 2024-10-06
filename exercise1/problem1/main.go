package main

import "fmt"

func main() {
	fmt.Println(addUp(600))
}

func addUp(number int) int {
	total := 0
	for i := 1; i <= number; i++ {
		total += i
	}
	return total
}
