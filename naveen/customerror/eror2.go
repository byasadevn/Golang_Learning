package errors

type errorCode struct {
	s string
}

func (e *errorCode) Error() string {
	return e.s
}

func New(text string) error {
	return &errorCode{text}
}
