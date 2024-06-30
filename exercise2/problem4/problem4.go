package problem4

type Element[T any] struct {
	value T
	next  *Element[T]
}

type LinkedList[T any] struct {
	list []Element[T]
	size int
}

func (ll *LinkedList[T]) Add(item *Element[T]) {}

func (ll *LinkedList[T]) Insert(item *Element[T], position int) error {
	return nil
}

func (ll *LinkedList[T]) Delete(item *Element[T]) error {
	return nil
}

func (ll *LinkedList[T]) Find(item any) (*Element[T], error) {
	return nil, nil
}

func (ll *LinkedList[T]) List() []Element[T] {
	return ll.list
}

func (ll *LinkedList[T]) Size() int {
	return ll.size
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}
