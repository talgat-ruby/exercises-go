package main

import "fmt"

func main() {
	var a = addUp(4)
	var b = addUp(13)
	var c = addUp(600)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func addUp(num int) int {
	return (num * (num + 1) / 2)
}
