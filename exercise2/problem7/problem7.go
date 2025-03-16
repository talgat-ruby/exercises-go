package main

import "fmt"

// swap takes pointers to two integers and swaps their values
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b) // Output: 2, 1
}
