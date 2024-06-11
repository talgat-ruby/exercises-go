package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func sumOfAllLegsNum(creatures ...interface{}) int {
	sum := 0
	for _, creature := range creatures {
		switch c := creature.(type) {
		case *Animal:
			sum += c.legsNum
		case *Insect:
			sum += c.legsNum
		}
	}
	return sum
}
