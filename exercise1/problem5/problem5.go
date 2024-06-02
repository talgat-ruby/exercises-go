package problem5

import "sort"

func products(productsMap map[string]int, price int) []string {
	var filteredProducts []string
	for product, cost := range productsMap {
		if cost >= price {
			filteredProducts = append(filteredProducts, product)
		}
	}

	sort.Slice(filteredProducts, func(i, j int) bool {
		if productsMap[filteredProducts[i]] == productsMap[filteredProducts[j]] {
			return filteredProducts[i] > filteredProducts[j]
		}
		return productsMap[filteredProducts[i]] < productsMap[filteredProducts[j]]
	})

	return filteredProducts
}
