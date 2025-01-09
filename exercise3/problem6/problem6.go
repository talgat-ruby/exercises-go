package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

type Legs interface {
	getLegsCount() int
}

func (animal *Animal) getLegsCount() int {
	return animal.legsNum
}

func (insect *Insect) getLegsCount() int {
	return insect.legsNum
}

func sumOfAllLegsNum(inputs ...Legs) int {
	total := 0
	for _, el := range inputs {
		total += el.getLegsCount()
	}
	return total
}
