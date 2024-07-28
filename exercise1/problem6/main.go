package main

import "strings"

func emojify(str string) string {
	emojiMap := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	words := strings.Fields(str)
	for i, word := range words {
		cleanWord := strings.Trim(word, ",.!?")
		if emoji, ok := emojiMap[cleanWord]; ok {
			words[i] = strings.Replace(word, cleanWord, emoji, 1)
		}
	}

	return strings.Join(words, " ")
}
