package problem5

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) compareAge(anotherPerson *Person) string {
	if anotherPerson.age > p.age {
		return fmt.Sprintf("%s is older than me.", anotherPerson.name)
	} else if anotherPerson.age < p.age {
		return fmt.Sprintf("%s is younger than me.", anotherPerson.name)
	} else {
		return fmt.Sprintf("%s is the same age as me.", anotherPerson.name)
	}
}
