package problem6

type Legged interface {
	Legs() int
}

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

func sumOfAllLegsNum(withLegs ...Legged) int {
	legs := 0
	for _, animalWithLeg := range withLegs {
		legs += animalWithLeg.Legs()
	}
	return legs
}
