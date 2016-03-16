package cast

// Error is a trivial implementation of error.
type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

// NewCastError returns an error that formats as the given text
func NewCastError(text string) *Error {
	return &Error{text}
}
