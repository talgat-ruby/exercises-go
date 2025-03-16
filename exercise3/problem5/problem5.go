package problem5

import "fmt"

type Person struct {
	name string
	age  int
}

func (me *Person) compareAge(other *Person) string {
	if other.age > me.age {
		return fmt.Sprintf("%s is older than me.", other.name)
	}

	if other.age < me.age {
		return fmt.Sprintf("%s is younger than me.", other.name)
	}

	return fmt.Sprintf("%s is the same age as me.", other.name)
}
