package main

import "fmt"

// Person represents a person with a name and age
type Person struct {
	name string
	age  int
}

// compareAge compares the age of another person with the receiver's age and returns a formatted string
func (p *Person) compareAge(other *Person) string {
	if p.age < other.age {
		return fmt.Sprintf("%s is older than me.", other.name)
	} else if p.age > other.age {
		return fmt.Sprintf("%s is younger than me.", other.name)
	}
	return fmt.Sprintf("%s is the same age as me.", other.name)
}

func main() {
	p1 := &Person{"Samuel", 24}
	p2 := &Person{"Joel", 36}
	p3 := &Person{"Lily", 24}

	fmt.Println(p1.compareAge(p2)) // Output: "Joel is older than me."
	fmt.Println(p2.compareAge(p1)) // Output: "Samuel is younger than me."
	fmt.Println(p1.compareAge(p3)) // Output: "Lily is the same age as me."
}
