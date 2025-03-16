package main

import "strings"

func emojify(s string) string {
	replacements := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	words := strings.Fields(s)
	for i, word := range words {
		cleanedWord := strings.Trim(word, ",.!?")

		if emoji, exists := replacements[cleanedWord]; exists {
			words[i] = strings.ReplaceAll(word, cleanedWord, emoji)
		}
	}

	return strings.Join(words, " ")
}
