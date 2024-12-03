package problem6

type Animal struct {
	name    string
	legsNum int
}

func (a *Animal) GetLegs() int {
	return a.legsNum
}

type Insect struct {
	name    string
	legsNum int
}

func (i *Insect) GetLegs() int {
	return i.legsNum
}

type HasLegs interface {
	GetLegs() int
}

func sumOfAllLegsNum(list ...HasLegs) int {
	total := 0
	for _, l := range list {
		total += l.GetLegs()
	}
	return total
}
