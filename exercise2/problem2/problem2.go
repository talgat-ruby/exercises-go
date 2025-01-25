package problem2

import "strings"

func capitalize(names []string) []string {
	capital := []string{}
	names1 := names

	for _, v := range names1 {
		temp := ""
		for i, v1 := range v {
			if len(v) > 0 && v[0] >= 97 && v[0] <= 122 && i == 0 {

				newv1 := strings.ToUpper(string(v1))
				temp += newv1

				newv1 = ""
			} else if len(v) > 0 && v1 >= 65 && v1 <= 90 && i > 0 {

				newv2 := strings.ToLower(string(v1))
				temp += newv2
				newv2 = ""
			} else {
				temp += string(v1)
			}

		}
		capital = append(capital, temp)
		temp = ""
	}
	return capital
}
