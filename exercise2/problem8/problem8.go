package problem8

import "fmt"

func simplify(list []string) map[string]int {
	indMap := make(map[string]int)
	load(indMap, list)

	return indMap
}

func load(m map[string]int, students []string) {
	for i, name := range students {
		m[name] = i
	}
}

func main() {

	result1 := simplify([]string{"a", "b", "c"})
	fmt.Println(result1)

	result2 := simplify([]string{"z", "y", "x", "u", "v"})
	fmt.Println(result2)
}
