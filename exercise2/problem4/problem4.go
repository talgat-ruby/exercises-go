package problem4

import "strings"

func mapping(letters []string) map[string]string {
	mp := make(map[string]string)

	for _, l := range letters {
		mp[strings.ToLower(l)] = strings.ToUpper(l)
	}

	return mp
}
