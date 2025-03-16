package main

import (
	"fmt"
	"strings"
)

func countVowels(s string) int {
	vowels := "aeiou"
	count := 0
	for _, char := range s {
		lowerChar := strings.ToLower(string(char))
		if strings.ContainsRune(vowels, rune(lowerChar[0])) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countVowels("Celebration"))
	fmt.Println(countVowels("Palm"))
	fmt.Println(countVowels("Prediction"))
}
