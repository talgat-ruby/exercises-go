package statusError

type StatusError struct {
	num     int
	message string
}

func New(status int, msg string) error {
	return &StatusError{status, msg}
}

func (st *StatusError) Error() string {
	return st.message
}

func (st *StatusError) Status() int {
	return st.num
}
