package main

import (
	"fmt"
	"strings"
)

func potatoes(s string) int {
	count := 0
	substring := "potato"

	for {
		index := strings.Index(s, substring)
		if index == -1 {
			break
		}
		count++
		s = s[index+len(substring):]
	}

	return count
}

func main() {
	fmt.Println(potatoes("potato"))
	fmt.Println(potatoes("potatopotato"))
	fmt.Println(potatoes("potatoapple"))
	fmt.Println(potatoes("potatopotatopotat"))
}
