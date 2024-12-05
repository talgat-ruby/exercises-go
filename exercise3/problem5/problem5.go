package problem5

type Person struct {
	name string
	age  int
}

func (person1 *Person) compareAge(person2 *Person) string {
	var answer string
	answer = person2.name + " is the same age as me."
	if person1.age < person2.age {
		answer = person2.name + " is older than me."
	} else if person1.age > person2.age {
		answer = person2.name + " is younger than me."
	}
	return answer
}
