package main

func countVowels(word string) int {
	vowels := []string{"a", "e", "i", "o", "u"}

	counter := 0
	for _, w := range word {
		for _, vowel := range vowels {
			if string(w) == vowel {
				counter++
			}
		}
	}

	return counter
}
