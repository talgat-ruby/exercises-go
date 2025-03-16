package main

import "fmt"

// Legged represents any entity with legs that can return a leg count
type Legged interface {
	GetLegsNum() int
}

// Animal represents an animal with a name and a number of legs
type Animal struct {
	name    string
	legsNum int
}

// GetLegsNum returns the number of legs for an Animal
func (a *Animal) GetLegsNum() int {
	return a.legsNum
}

// Insect represents an insect with a name and a number of legs
type Insect struct {
	name    string
	legsNum int
}

// GetLegsNum returns the number of legs for an Insect
func (i *Insect) GetLegsNum() int {
	return i.legsNum
}

// sumOfAllLegsNum calculates the total number of legs for all provided Legged entities
func sumOfAllLegsNum(leggeds ...Legged) int {
	totalLegs := 0
	for _, legged := range leggeds {
		totalLegs += legged.GetLegsNum()
	}
	return totalLegs
}

func main() {
	horse := &Animal{name: "horse", legsNum: 4}
	kangaroo := &Animal{name: "kangaroo", legsNum: 2}
	ant := &Insect{name: "ant", legsNum: 6}
	spider := &Insect{name: "spider", legsNum: 8}

	fmt.Println(sumOfAllLegsNum(horse, kangaroo, ant, spider)) // Output: 20
}
