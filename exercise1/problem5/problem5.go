package problem5

import (
	"fmt"
	"sort"
)

type Product struct {
	Name  string
	Price int
}

func products(catalog map[string]int, minPrice int) []string {
	var filteredProducts []Product
	for name, price := range catalog {
		if price >= minPrice {
			filteredProducts = append(filteredProducts, Product{Name: name, Price: price})
		}
	}

	fmt.Println("Filtered Products before sorting: ", filteredProducts)

	// Сортировка продуктов по возрастанию цены и убыванию имени при одинаковой цене
	sort.Slice(filteredProducts, func(i, j int) bool {
		fmt.Println("Comparing:", filteredProducts[i], "and", filteredProducts[j]) // Debug print
		if filteredProducts[i].Price == filteredProducts[j].Price {
			return filteredProducts[i].Name < filteredProducts[j].Name
		}
		return filteredProducts[i].Price > filteredProducts[j].Price
	})

	fmt.Println("Filtered and Sorted Products: ", filteredProducts)

	// Создание списка имен продуктов в нужном порядке
	var result []string
	for _, product := range filteredProducts {
		result = append(result, product.Name)
	}

	return result

}
