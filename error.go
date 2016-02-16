package cast

// castError is a trivial implementation of error.
type castError struct {
	msg string
}

func (e *castError) Error() string {
	return e.msg
}

// NewCastError returns an error that formats as the given text
func NewCastError(text string) error {
	return &castError{text}
}
