package main

import "fmt"

func main() {
	//{666, "1010011010"},
	fmt.Println(binary(0))
}

func binary(num int) string {
	if num == 0 {
		return "0"
	}
	bin := ""
	for num > 0 {
		remainder := num % 2
		bin = string(rune(remainder)+'0') + bin
		num = num / 2
	}
	return bin
}
