package main

import (
	"strconv"
)

func binary(decimalStr string) (string, error) {
	decimal, err := strconv.Atoi(decimalStr)
	if err != nil {
		return "", err
	}

	// Convert the integer to a binary string
	binaryStr := strconv.FormatInt(int64(decimal), 2)
	return binaryStr, nil
}
