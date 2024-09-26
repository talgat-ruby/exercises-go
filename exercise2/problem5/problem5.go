package problem5

import "sort"

func products(catalog map[string]int, minPrice int) []string {

	var filteredProducts []string
	for product, price := range catalog {
		if price >= minPrice {
			filteredProducts = append(filteredProducts, product)
		}
	}

	sort.Slice(filteredProducts, func(i, j int) bool {

		if catalog[filteredProducts[i]] != catalog[filteredProducts[j]] {

			return catalog[filteredProducts[i]] > catalog[filteredProducts[j]]
		}

		return filteredProducts[i] < filteredProducts[j]
	})

	return filteredProducts
}
