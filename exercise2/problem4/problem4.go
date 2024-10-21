package problem4

import (
	"fmt"
	"strings"
)

func mapping(lowerCase []string) map[string]string {
	result := make(map[string]string)

	for _, s := range lowerCase {
		result[s] = strings.ToUpper(s)
	}
	return result
}
func main() {
	fmt.Println(mapping([]string{"p", "s"}))
}
