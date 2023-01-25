package berr

import (
	"strings"

	"github.com/rheisen/berr/berrconst"
)

type Errors []Error

func (e Errors) Error() string {
	if e == nil {
		return ""
	}

	builder := strings.Builder{}
	builder.WriteString("[")

	for idx, val := range e {
		if idx > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(val.Error())
	}

	builder.WriteString("]")

	return builder.String()
}

type Error interface {
	error
	ErrorType() berrconst.BerrType
	ErrorMessage() string
	ErrorDetail() map[string]any
	Map() map[string]any
}
