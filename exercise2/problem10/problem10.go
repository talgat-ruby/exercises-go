package problem10

import (
	"fmt"
)

func factory() (map[string]int, func(string) func(int) int) {
	brands := make(map[string]int)

	makeBrand := func(brand string) func(int) int {

		if _, exists := brands[brand]; !exists {
			brands[brand] = 0
		}

		return func(count int) int {

			brands[brand] += count
			return brands[brand]
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
	makeBrand("Kia")

	fmt.Println(brands)
}
