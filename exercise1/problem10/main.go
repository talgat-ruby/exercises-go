package main

import (
	"fmt"
	"strconv"
)

// Функция sum суммирует два числа, представленных в виде строк.
// Если одно из значений не является допустимым числом, возвращает ошибку.
func sum(a, b string) (string, error) {
	numA, err := strconv.Atoi(a)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}

	numB, err := strconv.Atoi(b)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	result := numA + numB
	return strconv.Itoa(result), nil
}
