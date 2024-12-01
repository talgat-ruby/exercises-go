package main

import (
	"fmt"
	"unicode"
)

// detectWord function extracts the lowercase word from the mixed crowd of letters
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
	// Примеры использования функции detectWord
	fmt.Println(detectWord("UcUNFYGaFYFYGtNUH"))                    // "cat"
	fmt.Println(detectWord("bEEFGBuFBRrHgUHlNFYaYr"))               // "burglar"
	fmt.Println(detectWord("YFemHUFBbezFBYzFBYLleGBYEFGBMENTment")) // "embezzlement"
}
