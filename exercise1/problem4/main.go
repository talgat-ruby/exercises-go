package main

import (
	"fmt"
	"unicode"
)

func detectWord(crowd string) string {
	var word []rune
	for _, char := range crowd {
		if unicode.IsLower(char) {
			word = append(word, char)
		}
	}
	return string(word)
}

func main() {
	fmt.Println(detectWord("UcUNFYGaFYFYGtNUH"))
	fmt.Println(detectWord("bEEFGBuFBRrHgUHlNFYaYr"))
	fmt.Println(detectWord("YFemHUFBbezFBYzFBYLleGBYEFGBMENTment"))
}
