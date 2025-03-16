package problem5

type Person struct {
	Name string
	Age  int
}

func (p *Person) compareAge(p2 *Person) string {
	var compare string = " is older than me."
	if p2.Age == p.Age {
		compare = " is the same age as me."
	}
	if p2.Age < p.Age {
		compare = " is younger than me."
	}
	return p2.Name + compare
}
