package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func sumOfAllLegsNum(entities ...interface{}) int {
	total := 0
	for _, entity := range entities {
		switch e := entity.(type) {
		case Animal:
			total += e.legsNum
		case Insect:
			total += e.legsNum
		}
	}
	return total
}
