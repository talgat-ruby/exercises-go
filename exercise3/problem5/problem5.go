package problem5

type Person struct {
	name string
	age  int
}

func (p *Person) compareAge(other *Person) string {
	if p.age < other.age {
		return other.name + " is older than me."
	} else if p.age > other.age {
		return other.name + " is younger than me."
	} else {
		return other.name + " is the same age as me."
	}
}
