package main

import (
	"errors"
)

func sum(left, right string) (string, error) {
	leftNum, err := stringToNumber(left)
	if err != nil {
		return "", err
	}
	rightNum, err := stringToNumber(right)
	if err != nil {
		return "", err
	}
	return numberToString(leftNum + rightNum), nil
}

func numberToString(num int) string {
	out := make([]rune, 0)
	for num > 0 {
		r := rune(num%10 + '0')
		out = append([]rune{r}, out...)
		num = num / 10
	}
	return string(out)
}

func stringToNumber(s string) (int, error) {
	out := 0
	length := len(s)
	for i, c := range s {
		n := runeToNumber(c)
		if n > 9 || n < 0 {
			return 0, errors.New("NaN")
		}
		out += n * powerTen(length-i-1)
	}
	return out, nil
}

func runeToNumber(r rune) int {
	return int(r) - '0'
}

func powerTen(pow int) int {
	out := 1
	for range pow {
		out *= 10
	}
	return out
}
