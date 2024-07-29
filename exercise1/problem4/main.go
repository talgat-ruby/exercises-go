package main

func detectWord(word string) string {
	res := ""
	if len(word) == 0 {
		return res
	}

	wordRunes := []rune(word)
	for i := 0; i < len(wordRunes); i++ {
		if wordRunes[i] >= 'a' && wordRunes[i] <= 'z' {
			res += string(wordRunes[i])
		}
	}
	return res
}
