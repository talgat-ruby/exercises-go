package problem5

import "sort"

func products(productList map[string]int, price int) []string {

	var listOfProducts []string
	for key, value := range productList {
		if value > price {
			listOfProducts = append(listOfProducts, key)
		}
	}
	sort.Strings(listOfProducts)

	return listOfProducts

}
