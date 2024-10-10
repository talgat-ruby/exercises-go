package problem10

import "fmt"

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)

	makeBrand := func(brand string) func(int) {
		return func(amount int) {
			brands[brand] += amount
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

	_ = makeBrand("Kia")

	fmt.Println(brands)
}
