package problem6

type HasLegs interface {
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

func (a *Animal) Legs() int {
	return a.legsNum
}

func (i *Insect) Legs() int {
	return i.legsNum
}

func sumOfAllLegsNum(hasLegs ...HasLegs) int {
	var total int
	for _, v := range hasLegs {
		total += v.Legs()
	}
	return total

}
