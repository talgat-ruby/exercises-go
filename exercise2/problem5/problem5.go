package problem5

func products(pricesProducts map[string]int, minPrice int) []string {
	var result []string
	for i, v := range pricesProducts {
		if v > minPrice {
			result = append(result, i)
		}
	}
	sortProducts(result, pricesProducts)
	return result
}
func sortProducts(products []string, prices map[string]int) []string {
	n := len(products)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if prices[products[j]] < prices[products[j+1]] ||
				(prices[products[j]] == prices[products[j+1]] && products[j] > products[j+1]) {
				products[j], products[j+1] = products[j+1], products[j]
			}
		}
	}
	return products
}
