package problem5

type Person struct {
	name string
	age  int
}

func (person Person) compareAge(comparePerson *Person) string {
	if comparePerson.age > person.age {
		return comparePerson.name + " is older than me."
	} else if comparePerson.age < person.age {
		return comparePerson.name + " is younger than me."
	} else {
		return comparePerson.name + " is the same age as me."
	}
}
