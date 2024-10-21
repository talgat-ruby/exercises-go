package main

func countVowels(word string) int {
	vowels := []string{"a", "e", "i", "o", "u"}
	count := 0
	for _, i := range word {
		for _, j := range vowels {
			if string(i) == j {
				count++
			}
		}
	}
	return count
}
