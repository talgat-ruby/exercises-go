package problem6

type Inter interface {
	getName() string
	getLegsCount() int
}

func sumOfAllLegsNum(animals ...Inter) int {
	total := 0

	for _, animal := range animals {
		total += animal.getLegsCount()
	}

	return total
}
