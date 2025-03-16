package main

import "fmt"

// factory returns a map to track brand counts and a function to create brand incrementers
func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)

	// makeBrand takes a brand name and returns an incrementer function
	makeBrand := func(brand string) func(int) {
		// Initialize the brand in the map if it doesn't exist
		if _, exists := brands[brand]; !exists {
			brands[brand] = 0
		}
		
		// Incrementer function for the specified brand
		return func(count int) {
			brands[brand] += count
		}
	}

	return brands, makeBrand
}

func main() {
	brands, makeBrand := factory()
	
	toyotaIncrementer := makeBrand("Toyota")
	toyotaIncrementer(1)
	toyotaIncrementer(2)
	
	hyundaiIncrementer := makeBrand("Hyundai")
	hyundaiIncrementer(5)
	
	_ = makeBrand("Kia") // Kia is added to the map with a count of 0

	fmt.Println(brands) // Output: map[Hyundai:5 Kia:0 Toyota:3]
}
