package main

import "strings"

func emojify(text string) string {
	res := text
	replacements := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	for word, emoji := range replacements {
		res = strings.ReplaceAll(res, word, emoji)
	}

	return res
}
