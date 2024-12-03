package problem2

import (
	"fmt"
	"strings"
)

func capitalize(names []string) []string {
	capitalizedNames := make([]string, len(names))
	for i, name := range names {

		capitalizedNames[i] = strings.Title(strings.ToLower(name))
	}
	return capitalizedNames
}

func main() {
	fmt.Println(capitalize([]string{"mavis", "senaida", "letty"}))                // ["Mavis", "Senaida", "Letty"]
	fmt.Println(capitalize([]string{"samuel", "MABELLE", "letitia", "meridith"})) // ["Samuel", "Mabelle", "Letitia", "Meridith"]
	fmt.Println(capitalize([]string{"Slyvia", "Kristal", "Sharilyn", "Calista"})) // ["Slyvia", "Kristal", "Sharilyn", "Calista"]
}
