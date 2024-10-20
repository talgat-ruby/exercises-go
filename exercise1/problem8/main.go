package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countVowels("Celebration")) // 5
	fmt.Println(countVowels("Palm"))        // 1
	fmt.Println(countVowels("Prediction"))  // 4
}

func countVowels(word string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range word {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}
