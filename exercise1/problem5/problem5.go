package problem5

import (
	"sort"
)


func products(catalog map[string]int, minPrice int) []string {
	newCatalog := []string{}
	keys := []string{}

	for index := range catalog {
		keys = append(keys, index)
	}

	sort.Slice(keys, func(i, j int) bool {
		if catalog[keys[i]] == catalog[keys[j]] {
			return keys[i] < keys[j]
		}

		return catalog[keys[i]] > catalog[keys[j]] 
	})
	
	for _, value := range keys {
		if minPrice > catalog[value]{
			continue
		}
		
		newCatalog = append(newCatalog, value)
	}

	return newCatalog
}
