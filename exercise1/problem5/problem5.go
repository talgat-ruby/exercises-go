package problem5

import "sort"

type Product struct {
	name  string
	price int
}

func products(vals map[string]int, threshold int) []string {
	var filteredProducts []Product

	// Filter products based on the minimum price
	for name, price := range vals {
		if price >= threshold {
			filteredProducts = append(filteredProducts, Product{name, price})
		}
	}

	// Sort products by price in ascending order and by name in descending order for equal prices
	sort.Slice(filteredProducts, func(i, j int) bool {
		if filteredProducts[i].price == filteredProducts[j].price {
			return filteredProducts[i].name < filteredProducts[j].name // Descending order of names
		}
		return filteredProducts[i].price > filteredProducts[j].price // Ascending order of prices
	})

	// Extract the sorted product names
	var sortedProductNames []string
	for _, p := range filteredProducts {
		sortedProductNames = append(sortedProductNames, p.name)
	}

	return sortedProductNames
}
