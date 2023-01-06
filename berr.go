package berr

type Errors []Error

type Error interface {
	error
	ErrorType() string
	ErrorMessage() string
	ErrorDetail() map[string]any
	Map() map[string]any
}
