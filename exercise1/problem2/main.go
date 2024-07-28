package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(binary(1))
	fmt.Println(binary(5))
	fmt.Println(binary(10))
}

func binary(number int) string {
	binaryString := strconv.FormatInt(int64(number), 2)
	return binaryString
}
