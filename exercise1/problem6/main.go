package main

import "strings"

func emojify(a string) string {
	slovavsmile := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}
	for i, slovo := range slovavsmile {
		a = strings.ReplaceAll(a, i, slovo)
	}

	return a
}
