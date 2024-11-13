package main

import (
	"fmt"
	"strings"
)

// mapping takes a slice of lowercase letters and returns a map where each key is lowercase
// and the corresponding value is the uppercase version of that letter
func mapping(letters []string) map[string]string {
	letterMap := make(map[string]string)

	for _, letter := range letters {
		// Convert the letter to uppercase and add it to the map
		letterMap[letter] = strings.ToUpper(letter)
	}

	return letterMap
}

func main() {
	fmt.Println(mapping([]string{"p", "s"}))           // Output: { "p": "P", "s": "S" }
	fmt.Println(mapping([]string{"a", "b", "c"}))       // Output: { "a": "A", "b": "B", "c": "C" }
	fmt.Println(mapping([]string{"a", "v", "y", "z"}))  // Output: { "a": "A", "v": "V", "y": "Y", "z": "Z" }
}
