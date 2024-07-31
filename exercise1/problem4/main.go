package main

import "unicode"

func detectWord(mixedStr string) string {
    var hiddenWord []rune
    for _, char := range mixedStr {
        if unicode.IsLower(char) {
            hiddenWord = append(hiddenWord, char)
        }
    }
    return string(hiddenWord)
}
