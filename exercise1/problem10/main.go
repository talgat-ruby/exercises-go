package main

import (
	"fmt"
	"strconv"
)

func sum(str1 string, str2 string) (string, error) {
	val1, err := strconv.Atoi(string(str1))
	val2, err2 := strconv.Atoi(string(str2))

	if err != nil {
		return "", err
	} else if err2 != nil {
		return "", err2
	}

	return fmt.Sprintf("%d", val1+val2), nil
}
