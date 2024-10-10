package problem4

import (
	"fmt"
)

func mapping(letters []string) map[string]string {
	result := make(map[string]string)

	for _, letter := range letters {
		result[letter] = string(letter[0] - 32)
	}

	return result
}

func main() {
	fmt.Println(mapping([]string{"p", "s"}))
	fmt.Println(mapping([]string{"a", "b", "c"}))
	fmt.Println(mapping([]string{"a", "v", "y", "z"}))
}
