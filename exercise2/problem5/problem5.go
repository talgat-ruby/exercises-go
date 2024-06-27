package problem5

import "fmt"

type Person struct {
	name string
	age  int
}

func (p1 *Person) compareAge(p2 *Person) string {
	result := ""

	if p1.age > p2.age {
		result = "younger than"
	} else if p1.age < p2.age {
		result = "older than"
	} else if p1.age == p2.age {
		result = "the same age as"
	}

	return fmt.Sprintf("%s is %s me.", p2.name, result)
}
