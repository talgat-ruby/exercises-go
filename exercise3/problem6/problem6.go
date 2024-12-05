package problem6

// Legged interface defines a method to get the number of legs
type Legs interface {
	sumLegs() int
}

type Animal struct {
	name    string
	legsNum int
}

func (a *Animal) sumLegs() int {
	return a.legsNum
}

type Insect struct {
	name    string
	legsNum int
}

func (i *Insect) sumLegs() int {
	return i.legsNum
}

func sumOfAllLegsNum(l ...Legs) int {
	sum := 0
	for _, i := range l {
		sum += i.sumLegs()
	}
	return sum
}
