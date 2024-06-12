package problem5

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) compareAge(other *Person) string {
	if p.age < other.age {
		return fmt.Sprintf("%s is older than me.", other.name)
	} else if p.age > other.age {
		return fmt.Sprintf("%s is younger than me.", other.name)
	}
	return fmt.Sprintf("%s is the same age as me.", other.name)
}
