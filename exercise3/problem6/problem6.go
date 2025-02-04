package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func (a *Animal) Legs() int {
	return a.legsNum
}

func (i *Insect) Legs() int {
	return i.legsNum
}

type HasLegs interface {
	Legs() int
}

func sumOfAllLegsNum(legs ...HasLegs) int {
	var out int
	for _, leg := range legs {
		out += leg.Legs()
	}
	return out
}
