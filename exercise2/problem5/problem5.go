package main

import (
	"fmt"
	"sort"
)

// products filters and sorts product names based on the given minimum price
func products(items map[string]int, minPrice int) []string {
	// Create a slice to hold products that meet the price condition
	var filteredProducts []string

	// Filter products that have a price greater than or equal to minPrice
	for product, price := range items {
		if price >= minPrice {
			filteredProducts = append(filteredProducts, product)
		}
	}

	// Sort the products by price in descending order, and by name in ascending order for equal prices
	sort.Slice(filteredProducts, func(i, j int) bool {
		// Compare prices first
		if items[filteredProducts[i]] != items[filteredProducts[j]] {
			return items[filteredProducts[i]] > items[filteredProducts[j]]
		}
		// If prices are equal, sort by name in ascending order
		return filteredProducts[i] < filteredProducts[j]
	})

	return filteredProducts
}

func main() {
	fmt.Println(products(map[string]int{"Computer": 600, "TV": 800, "Radio": 50}, 300))     // Output: ["TV", "Computer"]
	fmt.Println(products(map[string]int{"Bike1": 510, "Bike2": 401, "Bike3": 501}, 500))     // Output: ["Bike1", "Bike3"]
	fmt.Println(products(map[string]int{"Loafers": 50, "Vans": 10, "Crocs": 20}, 100))       // Output: []
}
``
