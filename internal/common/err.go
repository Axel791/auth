package common

// AppError представляет ошибку с привязанным HTTP-кодом.
type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    400,
		Message: message,
	}
}

func NewInternalError(message string) *AppError {
	return &AppError{
		Code:    500,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    404,
		Message: message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    422,
		Message: message,
	}
}
