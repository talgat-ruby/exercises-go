package main

func emojify(sentence string) string {
	emojiMap := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	runes := []rune(sentence)
	result := ""
	i := 0

	for i < len(runes) {
		replaced := false
		for word, emoji := range emojiMap {
			wordRunes := []rune(word)
			if i+len(wordRunes) <= len(runes) && string(runes[i:i+len(wordRunes)]) == word {
				result += emoji
				i += len(wordRunes)
				replaced = true
				break
			}
		}
		if !replaced {
			result += string(runes[i])
			i++
		}
	}

	return result

}
