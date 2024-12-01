package main

import (
	"fmt"
	"strings"
)

// emojify function replaces specific words in a sentence with their corresponding emojis
func emojify(sentence string) string {
	// Define a map of words to emojis
	emojis := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	// Replace each word with its corresponding emoji
	for word, emoji := range emojis {
		sentence = strings.ReplaceAll(sentence, word, emoji)
	}

	return sentence
}

func main() {
	// Examples of using the emojify function
	fmt.Println(emojify("Make me smile")) // "Make me 🙂"
	fmt.Println(emojify("Make me grin"))  // "Make me 😀"
	fmt.Println(emojify("Make me sad"))   // "Make me 😥"
	fmt.Println(emojify("Make me mad"))   // "Make me 😠"
	fmt.Println(emojify("I am happy"))    // "I am happy" (no replacement)
}
