package main

import (
	"fmt"
	"strings"
)

func countVowels(s string) int {
	vowels := "aeiou"
	count := 0
	for _, char := range strings.ToLower(s) {
		if strings.ContainsRune(vowels, char) {
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
