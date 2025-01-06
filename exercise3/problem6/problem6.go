package problem6

type LegsProvider interface {
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

func sumOfAllLegsNum(entities ...LegsProvider) int {
	totalLegs := 0
	for _, entity := range entities {
		totalLegs += entity.GetLegsNum()
	}
	return totalLegs
}
