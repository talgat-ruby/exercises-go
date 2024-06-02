package problem2

func capitalize(words []string) []string {
	for i, word := range words {
		wordLength := len(word)
		if wordLength == 0 {
			continue
		}
		letters := []rune(word)
		for j := 0; j < wordLength; j++ {
			letter := letters[j]
			if 'a' <= letter && letter <= 'z' && j == 0 {
				letter = letter - 'a' + 'A'
			} else if 'A' <= letter && letter <= 'Z' && j != 0 {
				letter = letter - 'A' + 'a'
			}
			letters[j] = letter
		}
		words[i] = string(letters)
	}
	return words
}
