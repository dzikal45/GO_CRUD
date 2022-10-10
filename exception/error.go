package exception

type SomethingError struct {
	NotFoundError  string
	PasswordError  string
	somethingError string
}

func NewNotFoundError(error string) SomethingError {
	return SomethingError{NotFoundError: error}
}

func NewWrongPasswordError(error string) SomethingError {
	return SomethingError{PasswordError: error}
}
func NewFoundError(error string) SomethingError {
	return SomethingError{somethingError: error}
}
