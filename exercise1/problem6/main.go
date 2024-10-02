package main

import "strings"

func emojify(word string) string {
	if strings.Contains(word, "smile") {
		word = strings.ReplaceAll(word, "smile", "🙂")
	}
	if strings.Contains(word, "grin") {
		word = strings.ReplaceAll(word, "grin", "😀")
	}
	if strings.Contains(word, "sad") {
		word = strings.ReplaceAll(word, "sad", "😥")
	}
	if strings.Contains(word, "mad") {
		word = strings.ReplaceAll(word, "mad", "😠")
	}
	return word
}
