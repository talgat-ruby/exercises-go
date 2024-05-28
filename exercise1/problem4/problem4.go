package problem4

import "strings"

func mapping(in []string) map[string]string {
	result := map[string]string{}
	for _, e := range in {
		result[e] = strings.ToUpper(e)
	}
	return result
}
