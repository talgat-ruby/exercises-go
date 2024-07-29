package main

import (
	"strings"
)

func emojify(text string) string {
	// Заменяем ключевые слова на эмодзи
	text = strings.ReplaceAll(text, "smile", "🙂")
	text = strings.ReplaceAll(text, "grin", "😀")
	text = strings.ReplaceAll(text, "sad", "😥")
	text = strings.ReplaceAll(text, "mad", "😠")
	return text
}
