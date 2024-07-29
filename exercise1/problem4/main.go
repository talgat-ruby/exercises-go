package main

import (
	"fmt"
	"unicode"
)

func detectWord(word string) string {
	rr := []rune(word)
	var fw string
	for i := 0; i < len(word); i++ {
		if unicode.IsLower(rr[i]) {
			fw = fw + string(word[i])
			fmt.Printf("%c", word[i])
		}

	}
	return fw
}

func main() {
	var b string
	fmt.Scan(&b)
	detectWord(b)
}
