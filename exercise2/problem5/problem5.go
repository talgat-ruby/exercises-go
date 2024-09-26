package problem5

import (
	"fmt"
	"sort"
)

func products(productsMap map[string]int, minPrice int) []string {
	var filteredProducts []string

	for product, price := range productsMap {
		if price >= minPrice {
			filteredProducts = append(filteredProducts, product)
		}
	}

	sort.Slice(filteredProducts, func(i, j int) bool {
		if productsMap[filteredProducts[i]] == productsMap[filteredProducts[j]] {
			return filteredProducts[i] < filteredProducts[j]
		}
		return productsMap[filteredProducts[i]] > productsMap[filteredProducts[j]]
	})

	return filteredProducts
}

func main() {
	fmt.Println(products(map[string]int{"Computer": 600, "TV": 800, "Radio": 50}, 300))
	fmt.Println(products(map[string]int{"Bike1": 510, "Bike2": 401, "Bike3": 501}, 500))
	fmt.Println(products(map[string]int{"Loafers": 50, "Vans": 10, "Crocs": 20}, 100))
}
