package main

import (
	"strconv"
)

func binary(decimalStr string) (string, error) {
	decimal, err := strconv.Atoi(decimalStr)
	if err != nil {
		return "", err
	}

	binaryStr := strconv.FormatInt(int64(decimal), 2)
	return binaryStr, nil
}
