package main

import (
	"fmt"
	"strconv"
)

func binary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func main() {
	fmt.Println(binary(1))
	fmt.Println(binary(5))
	fmt.Println(binary(10))
}
