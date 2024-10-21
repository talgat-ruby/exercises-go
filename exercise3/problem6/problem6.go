package problem6

type Animal struct {
	name    string
	legsNum int
}

func (a *Animal) Legs() int {
	return a.legsNum
}

type Insect struct {
	name    string
	legsNum int
}

func (i *Insect) Legs() int {
	return i.legsNum
}

type NumbLegs interface {
	Legs() int
}

func sumOfAllLegsNum(names ...NumbLegs) int {
	sum := 0
	for _, n := range names {
		sum += n.Legs()
	}
	return sum
}
