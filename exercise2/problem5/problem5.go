package problem5

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) compareAge(person *Person) string {
	if p.Age < person.Age {
		return fmt.Sprintf("%s is older than me.", person.Name)
	} else if p.Age > person.Age {
		return fmt.Sprintf("%s is younger than me.", person.Name)
	}
	return fmt.Sprintf("%s is the same age as me.", person.Name)

}
