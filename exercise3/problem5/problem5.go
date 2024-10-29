package problem5

type Person struct {
	Name string
	Age  int
}

func (p *Person) compareAge(b *Person) string {
	if b.Age > p.Age {
		return b.Name + " is older than me."
	}

	if b.Age < p.Age {
		return b.Name + " is younger than me."
	}

	return b.Name + " is the same age as me."
}
