package main

import "strings"

func emojify(word string) string {

	matches := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	for k, v := range matches {
		word = strings.Replace(word, k, v, -1)
	}

	return word
}
