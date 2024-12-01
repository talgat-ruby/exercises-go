package problem6

type WithLegs interface {
	getLegsCount() int
}

type Animal struct {
	name    string
	legsNum int
}

func (a *Animal) getLegsCount() int {
	return a.legsNum
}

type Insect struct {
	name    string
	legsNum int
}

func (i *Insect) getLegsCount() int {
	return i.legsNum
}

func sumOfAllLegsNum(withLegs ...WithLegs) int {
	sum := 0
	for _, v := range withLegs {
		sum += v.getLegsCount()
	}
	return sum
}
