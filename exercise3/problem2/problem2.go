package problem2

import "errors"

type Stack struct {
	array []any
}

func (stack *Stack) Push(element any) {
	stack.array = append(stack.array, element)
}

func (stack *Stack) Pop() (any, error) {
	if stack.IsEmpty() {
		var zeroValue any
		return zeroValue, errors.New("queue is empty")
	}
	lastIndex := stack.Size() - 1
	val := stack.array[lastIndex]
	stack.array = stack.array[:lastIndex]
	return val, nil
}

func (stack *Stack) Peek() (any, error) {
	if stack.IsEmpty() {
		var zeroValue any
		return zeroValue, errors.New("queue is empty")
	}
	return stack.array[stack.Size()-1], nil
}

func (stack *Stack) Size() int {
	return len(stack.array)
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}
