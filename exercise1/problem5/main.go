package main

import (
	"fmt"
	"strings"
)

func potatoes(s string) int {
	return strings.Count(s, "potato")
}

func main() {
	fmt.Println(potatoes("potato"))
	fmt.Println(potatoes("potatopotato"))
	fmt.Println(potatoes("potatoapple"))
}
