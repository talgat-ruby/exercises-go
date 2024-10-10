package problem5

import (
	"fmt"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

// compareAge method to compare the age of two Person instances
func (p *Person) compareAge(other *Person) string {
	if p.Age < other.Age {
		return fmt.Sprintf("%s is older than me.", other.Name)
	} else if p.Age > other.Age {
		return fmt.Sprintf("%s is younger than me.", other.Name)
	}
	return fmt.Sprintf("%s is the same age as me.", other.Name)
}

func main() {
	p1 := &Person{"Samuel", 24}
	p2 := &Person{"Joel", 36}
	p3 := &Person{"Lily", 24}

	fmt.Println(p1.compareAge(p2)) // "Joel is older than me."
	fmt.Println(p2.compareAge(p1)) // "Samuel is younger than me."
	fmt.Println(p1.compareAge(p3)) // "Lily is the same age as me."
}
