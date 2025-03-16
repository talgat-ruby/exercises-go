package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func emojify(sentence string) string {
	emojiMap := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	words := strings.Fields(sentence)

	for i, word := range words {
		trimmedWord := strings.Trim(word, ",.!?")

		if newValue, exists := emojiMap[trimmedWord]; exists {
			words[i] = strings.Replace(word, trimmedWord, newValue, 1)
		}
	}

	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	result := emojify(input)

	fmt.Println(result)
}
