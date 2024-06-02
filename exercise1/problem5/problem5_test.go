package problem5

import (
	"slices"
	"testing"
)

func TestProducts(t *testing.T) {
	table := []struct {
		catalog  map[string]int
		minPrice int
		exp      []string
	}{
		{
			map[string]int{"Computer": 600, "TV": 800, "Radio": 100},
			300,
			[]string{"Computer", "TV"},
		},
		{
			map[string]int{"Bike1": 510, "Bike2": 401, "Bike3": 501},
			500,
			[]string{"Bike3", "Bike1"},
		},
		{
			map[string]int{"Calvin Klein": 2000, "Armani": 5000, "Dolce & Gabbana": 2000},
			1000,
			[]string{"Calvin Klein", "Dolce & Gabbana", "Armani"},
		},
		{
			map[string]int{"Loafers": 50, "Vans": 10, "Crocs": 20},
			100,
			[]string{},
		},
		{
			map[string]int{"Dell": 400, "HP": 300, "Apple": 400},
			350,
			[]string{"Apple", "Dell"},
		},
	}

	for _, r := range table {
		out := products(r.catalog, r.minPrice)
		if !slices.Equal(out, r.exp) {
			t.Errorf("products(%v, %d) was incorrect, got: %v, expected: %v.", r.catalog, r.minPrice, out, r.exp)
		}
	}
}
