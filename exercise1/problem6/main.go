package main

import "strings"

var emojiMap = map[string]string{
	"smile": "🙂",
	"grin":  "😀",
	"sad":   "😥",
	"mad":   "😠",
}

func emojify(sentence string) string {
	words := strings.Fields(sentence)
	for i, word := range words {

		cleanWord := strings.Trim(word, ".,!?;:")
		if emoji, ok := emojiMap[cleanWord]; ok {
			words[i] = strings.ReplaceAll(word, cleanWord, emoji)
		}
	}
	return strings.Join(words, " ")
}
