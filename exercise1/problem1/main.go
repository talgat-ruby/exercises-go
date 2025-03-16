package main

import "fmt"

func addUp(num int) int {
	return num * (num + 1) / 2
}

func main() {
	fmt.Println(addUp(4))
	fmt.Println(addUp(13))
	fmt.Println(addUp(600))
}
