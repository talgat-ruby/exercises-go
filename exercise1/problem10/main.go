package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	numA, errA := strconv.Atoi(a)
	if errA != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}
	numB, errB := strconv.Atoi(b)
	if errB != nil {
		return "", fmt.Errorf("string:%s cannot be converted", b)
	}
	result := numA + numB
	return strconv.Itoa(result), nil
}

func main() {
	fmt.Println(sum("1", "2"))
	fmt.Println(sum("10", "20"))
	fmt.Println(sum("a", "2"))
}
