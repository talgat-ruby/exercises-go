package problem6

type LegCounter interface {
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

func sumOfAllLegsNum(legCounters ...LegCounter) int {
	totalLegs := 0
	for _, lc := range legCounters {
		totalLegs += lc.Legs()
	}
	return totalLegs
}
