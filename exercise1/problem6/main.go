package main

import (
	"strings"
)

func emojify(sentence string) string {
	emojis := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	words := strings.Split(sentence, " ")

	for i, word := range words {
		trimmedWord := strings.TrimRight(word, ".,!?")
		if emoji, exists := emojis[strings.ToLower(trimmedWord)]; exists {
			if len(word) > len(trimmedWord) {
				words[i] = emoji + word[len(trimmedWord):]
			} else {
				words[i] = emoji
			}
		}
	}

	return strings.Join(words, " ")
}
