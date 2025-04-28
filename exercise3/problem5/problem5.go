package problem5

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) compareAge(person2 *Person) string {
	if person2.age > p.age {
		return fmt.Sprintf("%s is older than me.", person2.name)
	}

	if person2.age < p.age {
		return fmt.Sprintf("%s is younger than me.", person2.name)
	}

	return fmt.Sprintf("%s is the same age as me.", person2.name)
}
