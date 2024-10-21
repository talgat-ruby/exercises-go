// package main
//
// import "fmt"
//
//	func main() {
//		p1 := &Person{"Samuel", 24}
//		p2 := &Person{"Joel", 36}
//		p3 := &Person{"Lily", 24}
//
//		p1.compareAge(p2) // "Joel is older than me."
//		p2.compareAge(p1) // "Samuel is younger than me."
//		p1.compareAge(p3) // "Lily is the same age as me."
//	}
package problem5

type Person struct {
	Name string
	Age  int
}

func (p *Person) compareAge(person *Person) string {
	s := ""

	if person.Age > p.Age {
		s = "older than"
	} else if person.Age < p.Age {
		s = "younger than"
	} else {
		s = "the same age as"
	}

	return person.Name + " is " + s + " me."
}
