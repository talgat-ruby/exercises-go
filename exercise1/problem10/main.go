package main

import "fmt"

func sum(a, b string) (string, error) {
	if !isNumb(a) {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}
	if !isNumb(b) {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	numA := Atoi(a)
	numB := Atoi(b)

	result := Itoa(numA + numB)
	return result, nil
}

func isNumb(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if !(s[i] == '-' || s[i] == '+' || (s[i] >= '0' && s[i] <= '9')) {
				return false
			}
		} else {
			if !(s[i] >= '0' && s[i] <= '9') {
				return false
			}
		}
	}
	return true
}

func Atoi(s string) int {
	if !isNumb(s) {
		return 0
	}

	var res int
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}

	return res * sign
}

func Itoa(n int) string {
	res := ""
	if n == 0 {
		return "0"
	} else if n < 0 {
		n = -1 * n
		a := n
		for a > 0 {
			c := a%10 + 48
			a /= 10
			res = string(rune(c)) + res
		}
		res = "-" + res
	} else {
		a := n
		for a > 0 {
			c := a%10 + 48
			a /= 10
			res = string(rune(c)) + res
		}
	}

	return res
}
