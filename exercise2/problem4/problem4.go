package problem4

import "strings"

func mapping(arr []string) map[string]string {
	mapPair := make(map[string]string, len(arr))

	for _, v := range arr {
		mapPair[v] = strings.ToTitle(v)
	}
	return mapPair
}
