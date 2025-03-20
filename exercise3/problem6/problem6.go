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
	return horse.legsNum + kangaroo.legsNum + ant.legsNum + spider.legsNum
}
