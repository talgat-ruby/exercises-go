package problem5

import (
	"fmt"
)

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

func main() {
	p1 := &Person{"Samuel", 24}
	p2 := &Person{"Joel", 36}
	p3 := &Person{"Lily", 24}

	fmt.Println(p1.compareAge(p2)) 
	fmt.Println(p2.compareAge(p1)) 
	fmt.Println(p1.compareAge(p3)) 
}
