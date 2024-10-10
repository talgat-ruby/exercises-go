package problem2

import (
	"strings"
)

func capitalize(strs []string) []string {
	res := make([]string, 0, len(strs))
	for _, str := range strs {
		if str == "" {
			res = append(res, str)
			continue
		}
		if len(str) == 1 {
			res = append(res, strings.ToUpper(str))
			continue
		}

		newStr := strings.ToUpper(str[:1]) + strings.ToLower(str[1:])
		res = append(res, newStr)
	}
	return res
}

func main() {
	array1 := []string{"mavis", "senaida", "letty"}
	array2 := []string{"samuel", "MABELLE", "letitia", "meridith"}
	array3 := []string{"Slyvia", "Kristal", "Sharilyn", "Calista"}
	capitalize(array1)
	capitalize(array2)
	capitalize(array3)

}
