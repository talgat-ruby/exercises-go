package main

import (
    "fmt"
    "strings"
)

func emojify(sentence string) string {
    replacements := map[string]string{
        "smile": "🙂",
        "grin":  "😀",
        "sad":   "😥",
        "mad":   "😠",
    }

    words := strings.Fields(sentence)
    for i, word := range words {
        if emoji, ok := replacements[strings.ToLower(word)]; ok {
            words[i] = emoji
        }
    }

    return strings.Join(words, " ")
}

func main() {
    fmt.Println(emojify("Make me smile"))
    fmt.Println(emojify("Make me grin"))
    fmt.Println(emojify("Make me sad"))
    fmt.Println(emojify("I am mad at you"))
}