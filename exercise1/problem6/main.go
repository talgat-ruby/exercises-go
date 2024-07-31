package main

import (
	"fmt"
	"html"
	"strconv"
	"strings"
)

func main() {
	xx := "\\U0001f44f"

	// Hex String
	h := strings.ReplaceAll(xx, "\\U", "0x")

	// Hex to Int
	i, _ := strconv.ParseInt(h, 0, 64)

	// Unescape the string (HTML Entity -> String).
	str := html.UnescapeString(string(i))

	// Display the emoji.
	fmt.Println(str)
}
