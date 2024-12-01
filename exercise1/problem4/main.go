package main

import (
	"fmt"
	"unicode"
)

func detectWord(s string) string {
	var result []rune
	for _, char := range s {
		if unicode.IsLower(char) {
			result = append(result, char)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(detectWord("UcUNFYGaFYFYGtNUH"))
	fmt.Println(detectWord("bEEFGBuFBRrHgUHlNFYaYr"))
	fmt.Println(detectWord("YFemHUFBbezFBYzFBYLleGBYEFGBMENTment"))
}
