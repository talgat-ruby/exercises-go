package problem6

type creature interface {
	GetLegs() int
}

func (a *Animal) GetLegs() int {
	return a.legsNum
}

type Animal struct {
	name    string
	legsNum int
}

func (i *Insect) GetLegs() int {
	return i.legsNum
}

type Insect struct {
	name    string
	legsNum int
}

func sumOfAllLegsNum(s ...creature) int {
	var acc int = 0
	for _, v := range s {
		acc += v.GetLegs()
	}
	return acc
}
