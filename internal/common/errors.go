package common

// ErrorType represents different types of errors
type ErrorType string

const (
	ErrorTypeNotFound   ErrorType = "NOT_FOUND"
	ErrorTypeRateLimit  ErrorType = "RATE_LIMIT"
	ErrorTypeNoTokens   ErrorType = "NO_TOKENS"
	ErrorTypeMaxRetry   ErrorType = "MAX_RETRY"
	ErrorTypeBadRequest ErrorType = "BAD_REQUEST"
)

// CustomError represents a custom error with type
type CustomError struct {
	Message string
	Type    ErrorType
}

func (e *CustomError) Error() string {
	return e.Message
}

// NewCustomError creates a new custom error
func NewCustomError(message string, errorType ErrorType) *CustomError {
	return &CustomError{
		Message: message,
		Type:    errorType,
	}
}
