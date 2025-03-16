package problem5

type Person struct {
	Name string
	Age  int
}

func (p *Person) compareAge(value *Person) string {
	if p.Age < value.Age {
		return value.Name + " is older than me."
	} else if p.Age > value.Age {
		return value.Name + " is younger than me."
	} else {
		return value.Name + " is the same age as me."
	}
}
