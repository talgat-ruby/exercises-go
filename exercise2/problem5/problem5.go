package problem5

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) compareAge(p2 *Person) string {
	if p.age == p2.age {
		return fmt.Sprintf("%s is the same age as me.", p2.name)
	}

	if p.age > p2.age {
		return fmt.Sprintf("%s is younger than me.", p2.name)
	}

	return fmt.Sprintf("%s is older than me.", p2.name)
}
