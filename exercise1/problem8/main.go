package main

import (
	"fmt"
	"strings"
)

func countVowels(str string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countVowels("Celebration"))
}
