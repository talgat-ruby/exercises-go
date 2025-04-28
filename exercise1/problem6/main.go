package main

import "strings"

func emojify(sentence string) string {
	replacements := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	for word, emoji := range replacements {
		sentence = strings.ReplaceAll(sentence, word, emoji)
	}
	return sentence
}
