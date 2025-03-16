package main

import (
	"fmt"
	"strings"
)

func potatoes(a string) int {
	return strings.Count(a, "potato")
}

func main() {
	var c string
	fmt.Scan(&c)
	potatoes(c)
}
