package main

import "strings"

func emojify(words string) string {
	emojiMap := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}
	for k, v := range emojiMap {
		words = strings.ReplaceAll(words, k, v)
	}
	return words
}
