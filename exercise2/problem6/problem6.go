package problem6

type creature interface {
	Legs() int
}

type Animal struct {
	name    string
	legsNum int
}

func (a Animal) Legs() int {
	return a.legsNum
}

type Insect struct {
	name    string
	legsNum int
}

func (i Insect) Legs() int {
	return i.legsNum
}

func sumOfAllLegsNum(creatures ...creature) int {
	allLegsNum := 0
	for i := 0; i < len(creatures); i++ {
		allLegsNum += creatures[i].Legs()
	}
	return allLegsNum
}
