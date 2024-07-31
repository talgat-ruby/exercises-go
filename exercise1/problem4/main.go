package main

import (
	"fmt"
	"strings"
)

func detectWord(crowd string) string {
	var word strings.Builder
	for _, char := range crowd {
		if char >= 'a' && char <= 'z' {
			word.WriteRune(char)
		}
	}
	return word.String()
}

func main() {
	fmt.Println(detectWord("AQWEASDASDasdqwe"))
}
