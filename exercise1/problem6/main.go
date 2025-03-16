package main

import "strings"

var replacements map[string]rune = map[string]rune{
	"smile": '🙂',
	"grin":  '😀',
	"sad":   '😥',
	"mad":   '😠',
}

func emojify(crowd string) string {
	for word, emoji := range replacements {
		crowd = strings.ReplaceAll(crowd, word, string(emoji))
	}

	return crowd
}
