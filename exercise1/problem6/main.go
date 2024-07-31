package main

import (
	"fmt"
	"strings"
	"unicode"
)

func emojify(sentence string) string {
	emojiMap := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	isSeparator := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	words := strings.FieldsFunc(sentence, isSeparator)

	for _, word := range words {
		if emoji, exists := emojiMap[word]; exists {
			sentence = strings.Replace(sentence, word, emoji, 1)
		}
	}

	return sentence
}

func main() {
	fmt.Println(emojify("Make me smile"))
	fmt.Println(emojify("Make me grin"))
	fmt.Println(emojify("Make me sad"))
	fmt.Println(emojify("Make me mad"))
}
