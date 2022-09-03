package berrmodel

import (
	"fmt"

	"github.com/rheisen/berr/berrconst"
)

func NewBerr(errorType berrconst.BerrType, errorMessage string, errorDetail map[string]interface{}) *Berr {
	return &Berr{
		errorType:    errorType,
		errorMessage: errorMessage,
		errorDetail:  errorDetail,
	}
}

type Berr struct {
	errorType    berrconst.BerrType
	errorMessage string
	errorDetail  map[string]interface{}
}

func (e *Berr) Error() string {
	if e.errorDetail != nil && len(e.errorDetail) > 0 {
		return fmt.Sprintf(
			"%s: %s (%v)",
			e.errorType,
			e.errorMessage,
			e.errorDetail,
		)
	}

	return fmt.Sprintf(
		"%s: %s",
		e.errorType,
		e.errorMessage,
	)
}

func (e *Berr) ErrorType() berrconst.BerrType {
	return e.errorType
}

func (e *Berr) ErrorMessage() string {
	return e.errorMessage
}

func (e *Berr) ErrorDetail() map[string]interface{} {
	if e.errorDetail != nil && len(e.errorDetail) > 0 {
		detailCopy := make(map[string]interface{})
		for k, v := range e.errorDetail {
			detailCopy[k] = v
		}

		return detailCopy
	}

	return nil
}

func (e *Berr) Map() map[string]interface{} {
	return map[string]interface{}{
		"error_type": e.ErrorType().String(),
		"message":    e.ErrorMessage(),
		"detail":     e.ErrorDetail(),
	}
}
