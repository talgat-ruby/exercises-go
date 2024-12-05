package main

import (
	"strings"
)

func emojify(sentence string) string {

	emojiMap := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}
	words := strings.Fields(sentence)

	for i, word := range words {
		cleanedWord := strings.Trim(word, ".,!?")
		if emoji, found := emojiMap[cleanedWord]; found {
			words[i] = strings.Replace(word, cleanedWord, emoji, -1)
		}
	}

	return strings.Join(words, " ")
}
