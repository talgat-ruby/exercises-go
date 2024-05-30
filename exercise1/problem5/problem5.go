package problem5

import "sort"

type entry struct {
	key   string
	value int
}

func products(catalog map[string]int, minPrice int) []string {
	var preresults []entry
	var results []string
	for k, v := range catalog {
		if v > minPrice {
			preresults = append(preresults, entry{k, v})
		}
	}
	sort.Slice(preresults, func(a, b int) bool {
		if preresults[a].value == preresults[b].value {
			return preresults[a].key < preresults[b].key
		}
		return preresults[a].value > preresults[b].value
	})
	for i := 0; i < len(preresults); i++ {
		results = append(results, preresults[i].key)
	}
	return results
}
