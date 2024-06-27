package problem5

import (
	"sort"
)

type Product struct {
	product string
	price   int
}

func products(productList map[string]int, price int) []string {

	var listOfProducts []Product

	for oneProduct, productPrice := range productList {
		if price < productPrice {
			listOfProducts = append(listOfProducts, Product{oneProduct, productPrice})
		}
	}

	sort.SliceStable(listOfProducts, func(i, j int) bool {
		if listOfProducts[i].price == listOfProducts[j].price {
			return listOfProducts[i].product < listOfProducts[j].product
		}
		return listOfProducts[i].price > listOfProducts[j].price
	})

	var filteredProducts []string

	for _, product := range listOfProducts {
		filteredProducts = append(filteredProducts, product.product)
	}

	return filteredProducts
}
