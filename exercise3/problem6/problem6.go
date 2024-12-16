package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func sumOfAllLegsNum(legs ...interface{}) int {
	total := 0
	for _, leg := range legs {
		switch v := leg.(type) {
		case *Animal:
			total += v.legsNum
		case *Insect:
			total += v.legsNum
		}
	}
	return total
}
