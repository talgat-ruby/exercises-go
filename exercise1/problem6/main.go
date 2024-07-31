package main

import (
	"regexp"
	"strings"
)

func emojify(s string) string {
	emojiMap := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}
	re := regexp.MustCompile(`(\w+|[^\w\s]+|\s)`)
	words := re.FindAllString(s, -1)

	var result strings.Builder

	for _, word := range words {
		if emoji, exists := emojiMap[word]; exists {
			result.WriteString(emoji)
		} else {
			result.WriteString(word)
		}
	}
	return result.String()

}
