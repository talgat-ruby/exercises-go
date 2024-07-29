package main

import "fmt"
 
func addUp(a int) int {
	sum := 0
	for i := 1; i <= a; i++ {
		sum += i
	}
	return sum
}

func main() {
	var b int

	//fmt.Scan(&b)
	result := addUp(b)
	fmt.Println("addUp", result)
}
