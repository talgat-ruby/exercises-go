package problem7

func swap(firstNumber *int, secondNumber *int) (int, int) {

	var tempFirstNumber int
	tempFirstNumber = *firstNumber
	*firstNumber = *secondNumber
	*secondNumber = tempFirstNumber

	return *firstNumber, *secondNumber
}
