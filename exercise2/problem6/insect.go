package problem6

type Insect struct {
	name    string
	legsNum int
}

func (a Insect) getName() string {
	return a.name
}

func (a Insect) getLegsCount() int {
	return a.legsNum
}
