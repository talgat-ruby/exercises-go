package problem5

func products(catalog map[string]int, minPrice int) []string {
	var result []string
	for product, price := range catalog {
		if price >= minPrice {
			result = append(result, product)
		}
	}
	return result
}
