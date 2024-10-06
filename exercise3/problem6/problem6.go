package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func sumOfAllLegsNum(animals ...interface{}) int {
	totalLegsNum := 0
	for _, animal := range animals {
		switch a := animal.(type) {
		case *Animal:
			totalLegsNum += a.legsNum
		case *Insect:
			totalLegsNum += a.legsNum
		}
	}
	return totalLegsNum
}
