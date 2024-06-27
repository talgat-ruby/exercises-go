package problem6

type creature interface {
	Legs() int
}

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func (a Animal) Legs() int {
	return a.legsNum
}

func (i Insect) Legs() int {
	return i.legsNum
}

func sumOfAllLegsNum(creatures ...creature) int {
	total := 0

	for _, creature := range creatures {
		total += creature.Legs()
	}

	return total
}
