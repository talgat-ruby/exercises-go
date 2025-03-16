package main

import "strings"

func emojify(text string) string {
	text = strings.ReplaceAll(text, "smile", "🙂")
	text = strings.ReplaceAll(text, "grin", "😀")
	text = strings.ReplaceAll(text, "sad", "😥")
	text = strings.ReplaceAll(text, "mad", "😠")
	return text
}
