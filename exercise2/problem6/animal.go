package problem6

type Animal struct {
	name    string
	legsNum int
}

func (a Animal) getName() string {
	return a.name
}

func (a Animal) getLegsCount() int {
	return a.legsNum
}
