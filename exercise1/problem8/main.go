package main

import (
	"strings"
)

// Функция countVowels считает количество гласных букв в строке
func countVowels(text string) int {
	vowels := "aeiouAEIOU" // Множество гласных
	count := 0

	// Пройтись по каждому символу в строке
	for _, char := range text {
		// Проверить, является ли символ гласной
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
