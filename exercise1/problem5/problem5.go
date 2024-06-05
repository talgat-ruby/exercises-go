package problem5

func products(items map[string]int, minPrice int) []string {
	var minPriceProducts []string
	for product, amount := range items {
		if amount < minPrice {
			minPriceProducts = append(minPriceProducts, product)
		}
	}
	return minPriceProducts
}
