package problem6

type Animal struct{
	name string
	legsNum int
}

type Insect struct{
	name string
	legsNum int
}

func sumOfAllLegsNum(horse, kangaroo *Animal, ant, spider *Insect) int {
	count := 0

	count += horse.legsNum
	count += kangaroo.legsNum
	count += ant.legsNum
	count += spider.legsNum
	return count
}
