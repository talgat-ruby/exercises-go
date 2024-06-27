package problem5

type Person struct {
	name string
	age  int
}

func (p *Person) compareAge(s *Person) string {
	if p.age == s.age {
		return s.name + " is the same age as me."
	}
	if p.age > s.age {
		return s.name + " is younger than me."
	}
	if p.age < s.age {
		return s.name + " is older than me."
	}
	return ""
}
