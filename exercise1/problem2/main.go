package main

import "fmt"

func binary(n int) string {
	return fmt.Sprintf("%b", n)

}

func main() {
	fmt.Println(binary(1))
	fmt.Println(binary(5))
	fmt.Println(binary(10))
}
