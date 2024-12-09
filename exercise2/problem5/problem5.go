package problem5

import "sort"

type Catalog struct {
	Key   string
	Value int
}

type CatalogList []Catalog

func (c CatalogList) Len() int {
	return len(c)
}

func (c CatalogList) Less(i, j int) bool {
	if c[i].Value == c[j].Value {
		return c[i].Key < c[j].Key
	}
	return c[i].Value > c[j].Value
}

func (c CatalogList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func products(catalog map[string]int, minPrice int) []string {
	prices := make(CatalogList, len(catalog))
	i := 0

	for k, v := range catalog {
		if v > minPrice {
			prices[i] = Catalog{k, v}
			i++
		}
	}
	sort.Sort(prices)
	result := make([]string, 0)
	for _, v := range prices {
		if v.Key != "" {
			result = append(result, v.Key)
		}
	}
	return result
}
