package problem5

type Person struct {
	name string
	age  int
}

func (m *Person) compareAge(person *Person) string {
	if m.age < person.age {
		return person.name + " is older than me."
	} else if m.age > person.age {
		return person.name + " is younger than me."
	}
	return person.name + " is the same age as me."
}
