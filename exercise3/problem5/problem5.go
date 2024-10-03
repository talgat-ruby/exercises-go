package problem5

type Person struct {
	name string
	age  int
}

func (p *Person) compareAge(secondP *Person) string {
	if secondP.age > p.age {
		return secondP.name + " is older than me."
	} else if secondP.age < p.age {
		return secondP.name + " is younger than me."
	} else {
		return secondP.name + " is the same age as me."
	}
}
