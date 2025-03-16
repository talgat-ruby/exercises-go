package main

import (
    "strings"
)

func detectWord(crowd string) string {
    var word strings.Builder
    for _, ch := range crowd {
        if ch >= 'a' && ch <= 'z' {
            word.WriteRune(ch)
        }
    }
    return word.String()
}