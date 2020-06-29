package model

// Error struct contain error
type Error struct {
	Code    int
	Message string
	Debug   string
}

func (e *Error) Error() string {
	return e.Message
}

// NewError get custom error
func NewError(code int, message, debug string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Debug:   debug,
	}
}
