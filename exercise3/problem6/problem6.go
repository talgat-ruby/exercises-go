package problem6

type All interface {
	legCount() int
}

type Animal struct {
	name    string
	legsNum int
}

func (a *Animal) legCount() int {
	return a.legsNum
}

type Insect struct {
	name    string
	legsNum int
}

func (i *Insect) legCount() int {
	return i.legsNum
}

func sumOfAllLegsNum(legged ...All) int {
	sum := 0
	for _, a := range legged {
		sum += a.legCount()
	}
	return sum
}
