package problem6

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) CompareAge(other *Person) string {
	if p.Age > other.Age {
		return fmt.Sprintf("%s is younger than me.", other.Name)
	} else if p.Age < other.Age {
		return fmt.Sprintf("%s is older than me.", other.Name)
	} else {
		return fmt.Sprintf("%s is the same age as me.", other.Name)
	}
}

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func sumOfAllLegsNum(animals ...interface{}) int {
	totalLegs := 0
	for _, animal := range animals {
		switch a := animal.(type) {
		case *Animal:
			totalLegs += a.legsNum
		case *Insect:
			totalLegs += a.legsNum
		}
	}
	return totalLegs
}
