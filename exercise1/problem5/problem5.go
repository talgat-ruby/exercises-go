package problem5

import (
	"sort"
)

func products(productMap map[string]int, minPrice int) []string {
	length := len(productMap)
	out := make([]string, 0, length)
	invertedMap := make(map[int][]string)
	keys := make([]int, length)
	i := 0
	for k, v := range productMap {
		keys[i] = v
		i++
		invertedMap[v] = append(invertedMap[v], k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, k := range keys {
		if _, found := invertedMap[k]; found && k > minPrice {
			vals := invertedMap[k]
			sort.Strings(vals)
			out = append(out, vals...)
			delete(invertedMap, k)
		}
	}
	return out
}
