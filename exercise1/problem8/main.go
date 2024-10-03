package main

func countVowels(word string) int {
	res := 0

	if len(word) == 0 {
		return res
	}
	for i := 0; i < len(word); i++ {
		if word[i] == 'a' || word[i] == 'e' || word[i] == 'i' || word[i] == 'o' || word[i] == 'u' || word[i] == 'A' || word[i] == 'E' || word[i] == 'I' || word[i] == 'O' || word[i] == 'U' {
			res += 1
		}
	}
	return res
}
