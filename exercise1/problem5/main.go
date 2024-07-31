package main

import (
	"fmt"
	"strings"
)

func potatoes(a string) int {
	word := "potato"
	var b int
	b, _ = fmt.Println(strings.Count(a, word))
	return b
}

func main() {
	var c string
	fmt.Scan(&c)
	potatoes(c)
}
