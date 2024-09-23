package problem2

func capitalize(words []string) []string {
	out := make([]string, len(words))
	for i, word := range words {
		runeArray := make([]rune, len(word))
		for j, r := range word {
			if j == 0 && (r >= 'a' && r <= 'z') {
				r = r - 'a' + 'A'
			} else if j != 0 && (r >= 'A' && r <= 'Z') {
				r = r - 'A' + 'a'
			}
			runeArray[j] = r
		}
		out[i] = string(runeArray)
	}
	return out
}
