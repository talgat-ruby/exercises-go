package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func sumOfAllLegsNum(livingThings ...livingThing) int {
	result := 0
	for _, livingThing := range livingThings {
		result += livingThing.legsCount()
	}

	return result
}

type livingThing interface {
	legsCount() int
}

func (a *Animal) legsCount() int {
	return a.legsNum
}

func (i *Insect) legsCount() int {
	return i.legsNum
}
