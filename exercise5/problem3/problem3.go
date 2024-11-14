package problem3

import "fmt"

func sum(a, b int) int {
	c := make(chan int)

	go func(a, b int) {
		c <- a + b
	}(a, b)

	return <-c
}

func main() {
	c := sum(30, 40)
	fmt.Println("Sum:", c)
}
