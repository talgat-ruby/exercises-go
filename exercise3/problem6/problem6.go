package problem6

type Legged interface {
	GetLegsNum() int
}

type Animal struct {
	name    string
	legsNum int
}

func (a *Animal) GetLegsNum() int {
	return a.legsNum
}

type Insect struct {
	name    string
	legsNum int
}

func (i *Insect) GetLegsNum() int {
	return i.legsNum
}

func sumOfAllLegsNum(leggeds ...Legged) int {
	totalLegs := 0
	for _, legged := range leggeds {
		totalLegs += legged.GetLegsNum()
	}
	return totalLegs
}
