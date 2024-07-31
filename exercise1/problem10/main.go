package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(str1, str2 string) (string, error) {
	num1, err := strconv.Atoi(str1)
	if err != nil {
		return "", errors.New(fmt.Sprintf("string: %s cannot be converted", str1))
	}

	num2, err := strconv.Atoi(str2)
	if err != nil {
		return "", errors.New(fmt.Sprintf("string: %s cannot be converted", str2))
	}

	result := strconv.Itoa(num1 + num2)
	return result, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	res, err := sum(input1, input2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sum:", res)
	}
}
