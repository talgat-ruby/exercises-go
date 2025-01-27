package main

func emojify(sentence string) string {
	wordToEmoji := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	words := []rune(sentence)
	emojiSentence := ""

	i := 0
	for i < len(words) {
		match := false
		for word, emoji := range wordToEmoji {
			if i+len(word) <= len(words) && sentence[i:i+len(word)] == word {
				emojiSentence += emoji
				i += len(word)
				match = true
				break
			}
		}
		if !match {
			emojiSentence += string(words[i])
			i++
		}
	}

	return emojiSentence
}
