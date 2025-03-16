package main

import (
	"fmt"
	"strconv"
)

func binary(number int) string {
	return strconv.FormatInt(int64(number), 2)
}

func main() {
	fmt.Println(binary(1))
	fmt.Println(binary(5))
	fmt.Println(binary(10))
}
