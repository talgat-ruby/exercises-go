package main

import (
	"fmt"
	"strings"
)

// potatoes function returns the number of times "potato" appears in the input string
func potatoes(input string) int {
	return strings.Count(input, "potato")
}

func main() {
	// Примеры использования функции potatoes
	fmt.Println(potatoes("potato"))             // 1
	fmt.Println(potatoes("potatopotato"))       // 2
	fmt.Println(potatoes("potatoapple"))        // 1
	fmt.Println(potatoes("potatopotatopotato")) // 3
}
