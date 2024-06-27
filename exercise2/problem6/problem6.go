package problem6

type HasLegs interface {
	GetLegsNum() int
}

type Animal struct {
	name    string
	legsNum int
	HasLegs
}

type Insect struct {
	name    string
	legsNum int
	HasLegs
}

func (animal *Animal) GetLegsNum() int {
	return animal.legsNum
}

func (insect *Insect) GetLegsNum() int {
	return insect.legsNum
}

func sumOfAllLegsNum(livingThing ...interface{}) int {
	sum := 0
	for _, val := range livingThing {
		sum += val.(HasLegs).GetLegsNum()
	}
	return sum
}
