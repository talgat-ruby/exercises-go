package main

import "fmt"

func addUp(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(addUp(4))
	fmt.Println(addUp(13))
	fmt.Println(addUp(600))
}
