package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}
type LegsCounter interface {
	getLegsNum() int
}

func (a *Animal) getLegsNum() int {
	return a.legsNum
}
func (i *Insect) getLegsNum() int {
	return i.legsNum
}

func sumOfAllLegsNum(creatures ...LegsCounter) int {
	totalLegsNum := 0
	for _, creature := range creatures {
		totalLegsNum += creature.getLegsNum()
	}
	return totalLegsNum
}
