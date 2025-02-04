package problem4

func mapping(words []string) map[string]string {
	out := make(map[string]string)
	for _, word := range words {
		runes := []rune(word)
		runes[0] = runes[0] - 'a' + 'A'
		out[word] = string(runes)
	}
	return out
}
