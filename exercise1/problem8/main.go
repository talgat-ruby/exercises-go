package main

func countVowels(word string) int {
	count := 0
	for _, ch := range word {
		if ch == 'a' || ch == 'i' || ch == 'e' || ch == 'o' || ch == 'u' {
			count++
		}
	}
	return count
}
