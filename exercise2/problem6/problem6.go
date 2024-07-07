package problem6

type Leggable interface {
	GetLegs() int
}

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func (a Animal) GetLegs() int {
	return a.legsNum
}

func (a Insect) GetLegs() int {
	return a.legsNum
}

func sumOfAllLegsNum(creatures ...Leggable) int {
	sum := 0
	for _, creature := range creatures {
		sum += creature.GetLegs()
	}
	return sum
}
