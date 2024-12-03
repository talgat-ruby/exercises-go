package main

import (
	"fmt"
	"strings"
)

func detectWord(crowd string) string {
	var result strings.Builder
	for _, char := range crowd {
		if char >= 'a' && char <= 'z' {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	fmt.Println(detectWord("UcUNFYGaFYFYGtNUH"))
	fmt.Println(detectWord("bEEFGBuFBRrHgUHlNFYaYr"))
	fmt.Println(detectWord("YFemHUFBbezFBYzFBYLleGBYEFGBMENTment"))
}
