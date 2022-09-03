package berrinterface

import "github.com/rheisen/berr/berrconst"

type Berr interface {
	error
	ErrorType() berrconst.BerrType
	ErrorMessage() string
	ErrorDetail() map[string]interface{}
	Map() map[string]interface{}
}
