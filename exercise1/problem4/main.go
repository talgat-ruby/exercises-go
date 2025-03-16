package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(detectWord("UcUNFYGaFYFYGtNUH"))
	fmt.Println(detectWord("bEEFGBuFBRrHgUHlNFYaYr"))
	fmt.Println(detectWord("YFemHUFBbezFBYzFBYLleGBYEFGBMENTment"))
}

func detectWord(word string) string {
	var builder strings.Builder
	for _, r := range word {
		if !unicode.IsUpper(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
