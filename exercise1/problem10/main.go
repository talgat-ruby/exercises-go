package main

import (
	"fmt"
	"strconv"
)

func sum(x, y string) (string, error) {
	intX, errX := strconv.Atoi(x)
	if errX != nil {
		return "", fmt.Errorf("The string: %s cannot be converted", x)
	}

	intY, errY := strconv.Atoi(y)
	if errY != nil {
		return "", fmt.Errorf("The string: %s cannot be converted", y)
	}

	return strconv.Itoa(intX + intY), nil
}

func main() {
}
