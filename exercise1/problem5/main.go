package main

import (
	"fmt"
	"strings"
)

func potatoes(str string) int {
	return strings.Count(str, "potato")
}

func main() {
	fmt.Println(potatoes("potatopotatoasdq"))
}
