package problem5

type Person struct {
	name string
	age  int
}

func (p *Person) compareAge(o *Person) string {
	var conditionalAge string
	if o.age < p.age {
		conditionalAge = "younger than"
	} else if o.age > p.age {
		conditionalAge = "older than"
	} else {
		conditionalAge = "the same age as"
	}

	return o.name + " is " + conditionalAge + " me."
}
