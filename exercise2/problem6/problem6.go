package problem6

type All struct {
	name    string
	legsNum int
}

type Animal All

type Insect All

func sumOfAllLegsNum(animals ...interface{}) int {
	total := 0

	for _, animal := range animals {
		switch v := animal.(type) {
		case *Animal:
			total += v.legsNum
		case *Insect:
			total += v.legsNum
		}
	}

	return total
}
