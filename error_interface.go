package berr

import (
	"strings"
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
	Type() ErrorType
	Message() string
	Details() map[string]any
	String() string
	Map() map[string]any
	Unwrap() error
}
