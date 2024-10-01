package problem5

import "sort"

func products(a map[string]int, b int) []string {
	type Product struct {
		Name  string
		Price int
	}
	var list []Product
	for name, price := range a {
		if price >= b {
			list = append(list, Product{Name: name, Price: price})
		}
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].Price == list[j].Price {
			return list[i].Name < list[j].Name
		}
		return list[i].Price > list[j].Price
	})

	var res []string

	for _, i := range list {
		res = append(res, i.Name)
	}

	return res
}
