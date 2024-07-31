package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to compute bitwise AND
func bitwiseAND(a int, b int) int {
	result := a & b
	res1 := strconv.FormatInt(int64(result), 2)
	res2, _ := strconv.Atoi(res1)
	return res2
}

// Function to compute bitwise OR
func bitwiseOR(a int, b int) int {
	result := a | b
	res1 := strconv.FormatInt(int64(result), 2)
	res2, _ := strconv.Atoi(res1)
	return res2
}

// Function to compute bitwise XOR
func bitwiseXOR(a int, b int) int {
	result := a ^ b
	res1 := strconv.FormatInt(int64(result), 2)
	res2, _ := strconv.Atoi(res1)
	return res2
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the first number: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	num1, _ := strconv.Atoi(input1)

	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	num2, _ := strconv.Atoi(input2)

	andResult := bitwiseAND(num1, num2)
	orResult := bitwiseOR(num1, num2)
	xorResult := bitwiseXOR(num1, num2)

	fmt.Println(andResult)
	fmt.Println(orResult)
	fmt.Println(xorResult)

}
