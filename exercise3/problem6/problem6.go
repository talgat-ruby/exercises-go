package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

type HaveLeggs interface {
	getLeggs() int
}

func (animal *Animal) getLeggs() int {
	return animal.legsNum
}

func (insect *Insect) getLeggs() int {
	return insect.legsNum
}

func sumOfAllLegsNum(haveLeggs ...HaveLeggs) int {
	quantity := 0
	for _, value := range haveLeggs {
		quantity += value.getLeggs()
	}
	return quantity
}
