package main

import (
	"strings"
)

func potatoes(crowd string) int {
	return strings.Count(crowd, "potato")
}
