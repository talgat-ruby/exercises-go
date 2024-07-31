package main

import (
	"strings"
)

func emojify(sentence string) string {
	sentence = strings.ReplaceAll(sentence, "smile", "🙂")
	sentence = strings.ReplaceAll(sentence, "grin", "😀")
	sentence = strings.ReplaceAll(sentence, "sad", "😥")
	sentence = strings.ReplaceAll(sentence, "mad", "😠")

	return sentence
}

func main() {
}
