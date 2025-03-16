package main

import (
	"fmt"
	"unicode"
)

func detectWord(crowd string) string {
	var word string
	for _, char := range crowd {
		if unicode.IsLower(char) {
			word += string(char)
		}
	}
	return word
}

func main() {
	fmt.Println(detectWord("UcUNFYGaFYFYGtNUH"))
	fmt.Println(detectWord("bEEFGBuFBRrHgUHlNFYaYr"))
	fmt.Println(detectWord("YFemHUFBbezFBYzFBYLleGBYEFGBMENTment"))
}
