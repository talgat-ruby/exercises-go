package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}
type Legged interface {
	LegsCount() int
}

func (i Insect) LegsCount() int {
	return i.legsNum
}
func (a Animal) LegsCount() int {
	return a.legsNum
}

func sumOfAllLegsNum(creatures ...Legged) int {
	total := 0
	for _, creature := range creatures {
		total += creature.LegsCount()
	}
	return total
}
