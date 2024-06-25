package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func sumOfAllLegsNum(horse, kangaroo *Animal, ant, spider *Insect) int {
	sum := 0

	if horse != nil {
		sum += horse.legsNum
	}
	if kangaroo != nil {
		sum += kangaroo.legsNum
	}
	if ant != nil {
		sum += ant.legsNum
	}
	if spider != nil {
		sum += spider.legsNum
	}

	return sum
}
