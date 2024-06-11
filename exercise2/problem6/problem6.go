package problem6

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func sumOfAllLegsNum(animal1 *Animal, animal2 *Animal, insect1 *Insect, insect2 *Insect) int {
	return animal1.legsNum + animal2.legsNum + insect1.legsNum + insect2.legsNum
}
