package main

func countVowels(word string) int {
	count := 0
	vowels := map[rune]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}}
	for _, r := range word {
		if _, ok := vowels[r]; ok {
			count++
		}
	}
	return count
}
