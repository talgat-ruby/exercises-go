package problem5

import "sort"

func products(catalog map[string]int, minPrice int) []string {

	res := make([]string, 0, len(catalog))
	for product, price := range catalog {
		if price >= minPrice {
			res = append(res, product)
		}
	}

	sort.Strings(res)
	sort.Slice(res, func(i, j int) bool {
		return catalog[res[i]] > catalog[res[j]]
	})
	return res
}
