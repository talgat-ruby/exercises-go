package problem5

import "sort"

func products(catalog map[string]int, minPrice int) []string {
	var res []string
	for key, value := range catalog {
		if value > minPrice {
			res = append(res, key)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		if catalog[res[i]] == catalog[res[j]] {
			return res[i] < res[j]
		}

		return catalog[res[i]] > catalog[res[j]]
	})

	return res
}
