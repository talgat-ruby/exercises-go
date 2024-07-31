package main

import (
    "fmt"
)

func detectWord(crowd string) string {
    word := ""
    for _, letter := range crowd {
        if letter >= 'a' && letter <= 'z' {
            word += string(letter)
        }
    }
    return word
}

func main() {
    fmt.Println(detectWord("UcUNFYGaFYFYGtNUH"))
    fmt.Println(detectWord("bEEFGBuFBRrHgUHlNFYaYr"))
    fmt.Println(detectWord("YFemHUFBbezFBYzFBYLleGBYEFGBMENTment"))
}