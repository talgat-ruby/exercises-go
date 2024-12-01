package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(potatoes("potato"))       // 1
	fmt.Println(potatoes("potatopotato")) // 2
	fmt.Println(potatoes("potatoapple"))  // 1
}

func potatoes(s string) int {
	return strings.Count(s, "potato")
}
