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
	total := 0
	for _, creature := range creatures {
		switch v := creature.(type) {
		case *Animal:
			total += v.legsNum
		case *Insect:
			total += v.legsNum
		}
	}
	return total
}
