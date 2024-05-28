package problem5

func products(products map[string]int, minPrice int) []string {
	result := []string{}
	var price int
	for product, amount := range products {
		if amount >= minPrice && amount >= price {
			price = amount
			result = append(result, product)
		} else if amount > price {
			result = addElem(result, product)
		}
	}
	return result
}

func addElem(slice []string, elem string) []string {
	newSlice := make([]string, len(slice)+1)
	copy(newSlice[1:], slice)
	newSlice[0] = elem
	return newSlice
}
