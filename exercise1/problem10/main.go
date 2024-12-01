package main

import (
	"fmt"
	"strconv"
)

func sum(word1 string, word2 string) (string, error) {
	num1, err1 := strconv.Atoi(word1)
	if err1 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", word1)
	}

	num2, err2 := strconv.Atoi(word2)
	if err2 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", word2)
	}

	result := num1 + num2
	return strconv.Itoa(result), nil
}
