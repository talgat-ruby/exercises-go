package problem5

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) compareAge(per *Person) string {
	if p.age > per.age {
		return fmt.Sprintf("%s is younger than me.", per.name)
	} else if p.age < per.age {
		return fmt.Sprintf("%s is older than me.", per.name)
	}
	return fmt.Sprintf("%s is the same age as me.", per.name)
}
