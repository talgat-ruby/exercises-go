package problem5

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p1 Person) compareAge(p2 *Person) string {
	if p1.Age > p2.Age {
		return fmt.Sprintf("%s is younger than me.", p2.Name)
	}
	if p1.Age < p2.Age {
		return fmt.Sprintf("%s is older than me.", p2.Name)
	}
	return fmt.Sprintf("%s is the same age as me.", p2.Name)
}
