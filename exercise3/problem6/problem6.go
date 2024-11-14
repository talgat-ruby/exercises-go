package problem6

import "fmt"

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func main() {
	horse := &Animal{
		name:    "horse",
		legsNum: 4,
	}
	kangaroo := &Animal{
		name:    "kangaroo",
		legsNum: 2,
	}
	ant := &Insect{
		name:    "ant",
		legsNum: 6,
	}
	spider := &Animal{
		name:    "spider",
		legsNum: 8,
	}

	var sumOfAllLegs1 int = (horse.legsNum + kangaroo.legsNum)
	var sumOfAllLegs2 int = (ant.legsNum + spider.legsNum)

	var totalLegsNum int = (sumOfAllLegs1 + sumOfAllLegs2)

	fmt.Printf("sumOfAllLegsNum %d", totalLegsNum)
}
