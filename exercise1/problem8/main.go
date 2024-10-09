package main

func countVowels(word string) int {
	count := 0
	for _, ch := range word {
		if isVowel(ch) {
			count++
		}
	}
	return count
}
func isVowel(char rune) bool {
	if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
		char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
		return true
	}
	return false
}
