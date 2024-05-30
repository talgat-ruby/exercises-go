package problem5

func products(catalog map[string]int, price int) []string {
	products := make([]string, 0, len(catalog))
	for k, v := range catalog {
		if v < price {
			continue
		}

		products = append(products, k)
	}

	for i := 0; i < len(products)-1; i++ {
		for j := 0; j < len(products)-1-i; j++ {
			if catalog[products[j]] < catalog[products[j+1]] {
				products[j], products[j+1] = products[j+1], products[j]
			}

			if products[j] == products[j+1] && products[j] < products[j+1] {
				products[j], products[j+1] = products[j+1], products[j]
			}
		}
	}
	return products
}
