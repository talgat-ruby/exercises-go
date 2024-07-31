package main

import (
	"strings"
)

func emojify(res string) string {
	replacer := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}
	for word, emoji := range replacer {
		res = strings.ReplaceAll(res, word, emoji)
	}
	return res
}
