package structs

// Error represents a structured error type with additional context.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data   interface{} `json:"data"`
}

// NewError creates a new Error instance with the given code, message, and optional cause.
func NewError(code int, message string, data error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Data:   data,
	}
}
