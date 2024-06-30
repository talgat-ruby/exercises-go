package problem6

type Legged interface {
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

func sumOfAllLegsNum(livingBeings ...Legged) int {
	legs := 0
	for _, v := range livingBeings {
		legs += v.Legs()
	}
	return legs
}
