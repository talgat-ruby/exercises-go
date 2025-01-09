package main

func countVowels(s string) int {
	var vowels = []rune{'a', 'e', 'i', 'o', 'u'}
	count := 0
	for _, c := range s {
		for _, vowel := range vowels {
			if vowel == c {
				count++
			}
		}
	}
	return count
}
