package problem8

import (
	"fmt"
)

func simplify(s []string) map[string]int {
	result := make(map[string]int)

	for i, v := range s {
		result[v] = i
	}

	return result
}

func main() {
	fmt.Println(simplify([]string{"a", "b", "c"}))
	fmt.Println(simplify([]string{"z", "y", "x", "u", "v"}))
}
