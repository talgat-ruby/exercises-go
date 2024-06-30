package problem5

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) compareAge(otherPerson *Person) string {
	if otherPerson.age == p.age {
		return fmt.Sprintf("%s is the same age as me.", otherPerson.name)
	} else if otherPerson.age > p.age {
		return fmt.Sprintf("%s is older than me.", otherPerson.name)
	} else {
		return fmt.Sprintf("%s is younger than me.", otherPerson.name)
	}
}
