package problem6

type LegsNum interface {
	getLegsNum() int
}

type Animal struct {
	name    string
	legsNum int
}

func (animal Animal) getLegsNum() int {
	return animal.legsNum
}

type Insect struct {
	name    string
	legsNum int
}

func (insect Insect) getLegsNum() int {
	return insect.legsNum
}

func sumOfAllLegsNum(num ...LegsNum) int {
	var result int
	for _, n := range num {
		result += n.getLegsNum()
	}
	return result
}
