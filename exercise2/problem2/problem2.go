package main

import (
	"fmt"
	"strings"
)

// capitalize takes an array of names and returns an array where only the first letter of each name is capitalized
func capitalize(names []string) []string {
	var capitalizedNames []string

	for _, name := range names {
		// Capitalize the first letter and make the rest lowercase
		if len(name) > 0 {
			capitalized := strings.ToUpper(string(name[0])) + strings.ToLower(name[1:])
			capitalizedNames = append(capitalizedNames, capitalized)
		} else {
			// Handle any empty strings in the array
			capitalizedNames = append(capitalizedNames, "")
		}
	}

	return capitalizedNames
}

func main() {
	fmt.Println(capitalize([]string{"mavis", "senaida", "letty"}))      // Output: ["Mavis", "Senaida", "Letty"]
	fmt.Println(capitalize([]string{"samuel", "MABELLE", "letitia", "meridith"})) // Output: ["Samuel", "Mabelle", "Letitia", "Meridith"]
	fmt.Println(capitalize([]string{"Slyvia", "Kristal", "Sharilyn", "Calista"})) // Output: ["Slyvia", "Kristal", "Sharilyn", "Calista"]
}
