package problem4

import "strings"

func mapping(inpt []string) map[string]string {
	result := make(map[string]string)

	for _, in := range inpt {
		result[in] = strings.ToUpper(in)
	}

	return result
}
