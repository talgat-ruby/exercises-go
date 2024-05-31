package problem5

import "sort"

type Product struct {
	Name  string
	Price int
}

func products(catalog map[string]int, minPrice int) []string {
	var products []Product

	for name, price := range catalog {
		if price >= minPrice {
			products = append(products, Product{name, price})
		}
	}

	sort.Slice(products, func(i, j int) bool {
		if products[i].Price == products[j].Price {
			return products[i].Name < products[j].Name
		}
		return products[i].Price > products[j].Price
	})

	result := make([]string, len(products))
	for i, product := range products {
		result[i] = product.Name
	}

	return result
}
