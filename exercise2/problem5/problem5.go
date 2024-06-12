package problem5

type Person struct {
	Name string
	Age  int
}

func (p *Person) compareAge(p2 *Person) string {
	if p.Age < p2.Age {
		return p2.Name + " is older than me."
	} else if p.Age > p2.Age {
		return p2.Name + " is younger than me."
	}
	return p2.Name + " is the same age as me."
}
