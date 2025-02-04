package main

import "strings"

var emoji = map[string]string{
	"smile": "🙂",
	"grin":  "😀",
	"sad":   "😥",
	"mad":   "😠",
}

func emojify(text string) string {
	for i, v := range emoji {
		if strings.Contains(text, i) {
			text = strings.Replace(text, i, v, 1)
		}
	}
	return text
}
