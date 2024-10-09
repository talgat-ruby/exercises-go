package main

import "errors"

func main() {
	// fmt.Println("Hello, 世界")
	// sum("10", "20") // "30", nil
	sum("10", "20")
}

func sum(str1 string, str2 string) (string, error) {
	if !isInt(str1) || !isInt(str2) {
		return "", errors.New("string: a cannot be converted")
	}
	num1 := atoi(str1)
	num2 := atoi(str2)
	result := itoa(num1 + num2)
	return result, nil
}

func atoi(s string) int {
	base := 1
	num := 0
	for i := len(s) - 1; i >= 0; i-- {
		num += int(s[i]-'0') * base
		base *= 10
	}
	return num
}

func itoa(n int) string {
	num := ""
	for n > 0 {
		num = string(rune(n%10+'0')) + num
		n /= 10
	}

	return num
}

func isInt(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
