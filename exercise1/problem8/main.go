package main

import "fmt"

func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		for j := 0; j < len(vowels); j++ {
			if char == vowels[j] {
				count++
				break
			}
		}
	}

	return count
}

func main() {
	fmt.Println(countVowels("Celebration")) // 5
	fmt.Println(countVowels("Palm"))        // 1
	fmt.Println(countVowels("Prediction"))  // 4
}
