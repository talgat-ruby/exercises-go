package problem5

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (person1 *Person) compareAge(person2 *Person) string {
	var format string
	switch {
	case person1.Age > person2.Age:
		format = "%s is younger than me."
	case person1.Age < person2.Age:
		format = "%s is older than me."
	default:
		format = "%s is the same age as me."
	}
	return fmt.Sprintf(format, person2.Name)
}