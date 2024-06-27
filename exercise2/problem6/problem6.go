package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

type HasLegs interface {
	getLegsNum() int
}

func (a Animal) getLegsNum() int {
	return a.legsNum
}

func (i Insect) getLegsNum() int {
	return i.legsNum
}

func sumOfAllLegsNum(arr ...HasLegs) int {
	sum := 0
	for _, t := range arr {
		sum += t.getLegsNum()
	}

	return sum
}
