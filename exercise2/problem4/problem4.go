package problem4

import (
	"fmt"
	"strings"
)

func mapping(letters []string) map[string]string {
	result := make(map[string]string)
	for _, letter := range letters {

		result[letter] = strings.ToUpper(letter)
	}
	return result
}

func main() {

	fmt.Println(mapping([]string{"p", "s"}))
	fmt.Println(mapping([]string{"a", "b", "c"}))
	fmt.Println(mapping([]string{"a", "v", "y", "z"}))
}
