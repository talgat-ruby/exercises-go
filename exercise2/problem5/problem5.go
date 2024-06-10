package problem5

type Person struct {
	name string
	age  int
}

func (p *Person) compareAge(person *Person) string {
	if person.age == p.age {
		return person.name + " is the same age as me."
	} else if person.age > p.age {
		return person.name + " is older than me."
	} else {
		return person.name + " is younger than me."
	}
}
