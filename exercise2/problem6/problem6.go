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
		switch v := creature.(type) {
		case *Animal:
			sum += v.legsNum
		case *Insect:
			sum += v.legsNum
		}
	}
	return sum
}
