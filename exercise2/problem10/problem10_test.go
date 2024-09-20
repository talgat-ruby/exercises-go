package problem10

import (
	"maps"
	"testing"
)

func TestFactory(t *testing.T) {
	brands, makeBrand := factory()
	toyotaIncrementer := makeBrand("Toyota")
	toyotaIncrementer(1)
	toyotaIncrementer(2)
	hyundaiIncrementer := makeBrand("Hyundai")
	hyundaiIncrementer(5)
	makeBrand("Kia")

	exp := map[string]int{"Hyundai": 5, "Kia": 0, "Toyota": 3}

	if !maps.Equal(brands, exp) {
		t.Errorf("factory was incorrect, got: %v, expected: %v.", brands, exp)
	}
}
