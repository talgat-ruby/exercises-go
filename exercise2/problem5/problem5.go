package problem5

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p1 *Person) compareAge(p2 *Person) string {
	format := ""
	switch {
	case p1.Age > p2.Age:
		format = "%s is younger than me."
	case p1.Age < p2.Age:
		format = "%s is older than me."
	default:
		format = "%s is the same age as me."
	}
	return fmt.Sprintf(format, p2.Name)
}
