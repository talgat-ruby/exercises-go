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
	c := 0

	c += horse.legsNum
	c += kangaroo.legsNum
	c += ant.legsNum
	c += spider.legsNum
	return c
}
