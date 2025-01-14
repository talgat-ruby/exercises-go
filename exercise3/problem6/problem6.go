package problem6

type Len interface {
	Lens() int
}

type Animal struct {
	name    string
	legsNum int
}

func (a *Animal) Lens() int {
	return a.legsNum
}

type Insect struct {
	name    string
	legsNum int
}

func (i *Insect) Lens() int {
	return i.legsNum
}

func sumOfAllLegsNum(lens ...Len) int {
	sum := 0
	for _, v := range lens {
		sum = sum + v.Lens()
	}
	return sum
}
