package problem4

func mapping(words []string) map[string]string {
	out := make(map[string]string)
	for _, word := range words {
		out[word] = ToUpper(word)
	}
	return out
}

func ToUpper(word string) string {
	letters := []rune(word)
	upperWord := make([]rune, len(word))
	for i, letter := range letters {
		if letter >= 'a' && letter <= 'z' {
			upperWord[i] = letter - 'a' + 'A'
		}
	}
	return string(upperWord)
}
