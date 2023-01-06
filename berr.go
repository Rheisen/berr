package berr

import "github.com/rheisen/berr/berrconst"

type Errors []Error

type Error interface {
	error
	ErrorType() berrconst.BerrType
	ErrorMessage() string
	ErrorDetail() map[string]any
	Map() map[string]any
}
