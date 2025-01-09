package problem5

import "fmt"

type Person struct {
	name string
	age  int
}

func (person Person) compareAge(p *Person) string {
	if person.age < p.age {
		return fmt.Sprintf("%s is older than me.", p.name)
	}

	if person.age > p.age {
		return fmt.Sprintf("%s is younger than me.", p.name)
	}

	return fmt.Sprintf("%s is the same age as me.", p.name)
}
