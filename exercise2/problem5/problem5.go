package problem5

type Person struct {
	name string
	age  int
}

func (p Person) compareAge(otherPerson *Person) string {
	if p.age == otherPerson.age {
		return otherPerson.name + " is the same age as me."
	} else if p.age < otherPerson.age {
		return otherPerson.name + " is older than me."
	} else {
		return otherPerson.name + " is younger than me."
	}
}
