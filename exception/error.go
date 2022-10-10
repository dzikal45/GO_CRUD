package exception

type SomethingError struct {
	NotFoundError string
	PasswordError string
	BookIsBooked  string
}

func NewNotFoundError(error string) SomethingError {
	return SomethingError{NotFoundError: error}
}

func NewWrongPasswordError(error string) SomethingError {
	return SomethingError{PasswordError: error}
}
func NewBookIsBooked(error string) SomethingError {
	return SomethingError{BookIsBooked: error}
}
