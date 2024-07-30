package main

import (
	"fmt"
	"unicode"
)

func detectWord(s string) string {
	result := ""
	for _, char := range s {
		if unicode.IsLower(char) {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(detectWord("UcUNFYGaFYFYGtNUH"))
	fmt.Println(detectWord("bEEFGBuFBRrHgUHlNFYaYr"))
	fmt.Println(detectWord("YFemHUFBbezFBYzFBYLleGBYEFGBMENTment"))
}
