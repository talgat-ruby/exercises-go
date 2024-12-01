package problem5

import "sort"

func products(catalog map[string]int, minPrice int) (exp []string) {
	for product, price := range catalog {
		if price >= minPrice {
			exp = append(exp, product)
		}
	}
	sort.Slice(exp, func(i, j int) bool {
		if catalog[exp[i]] == catalog[exp[j]] {
			return exp[i] < exp[j]
		}
		return catalog[exp[i]] > catalog[exp[j]]
	})
	return
}
