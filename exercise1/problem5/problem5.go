package problem5

import "sort"

func products(catalog map[string]int, minPrice int) []string {
	var result []string
	keys := make([]string, 0, len(catalog))

	for key, value := range catalog {
		if value >= minPrice {
			keys = append(keys, key)
		}
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if catalog[keys[i]] > catalog[keys[j]] {
			return true
		} else if catalog[keys[i]] == catalog[keys[j]] {
			return keys[i] < keys[j]
		}
		return false
	})

	for k := range keys {
		result = append(result, keys[k])
	}

	return result
}
