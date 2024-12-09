package problem3

type counter struct {
	val int64
}

func newCounter() *counter {
	return &counter{
		val: 0,
	}
}
