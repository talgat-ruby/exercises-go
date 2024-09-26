package problem2

import (
	"fmt"
	"strings"
)

func capitalize(names []string) []string {
	capitalizedNames := make([]string, len(names))

	for i, name := range names {
		name = strings.ToLower(name)
		if len(name) > 0 {
			capitalizedNames[i] = strings.ToUpper(string(name[0])) + name[1:]
		} else {
			capitalizedNames[i] = name
		}
	}

	return capitalizedNames
}

func main() {
	fmt.Println(capitalize([]string{"mavis", "senaida", "letty"}))
	fmt.Println(capitalize([]string{"samuel", "MABELLE", "letitia", "meridith"}))
	fmt.Println(capitalize([]string{"Slyvia", "Kristal", "Sharilyn", "Calista"}))
}
