package problem5

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) compareAge(other *Person) string {
	if p.Age < other.Age {
		return fmt.Sprintf("%s is older than me.", other.Name)
	} else if p.Age > other.Age {
		return fmt.Sprintf("%s is younger than me.", other.Name)
	} else {
		return fmt.Sprintf("%s is the same age as me.", other.Name)
	}
}
