package berr

import (
	"fmt"

	"github.com/rheisen/berr/berrconst"
)

func newBerr(errorType berrconst.BerrType, errorMessage string, details ...D) *berr {
	errorDetail := make(map[string]any, len(details))
	for _, d := range details {
		errorDetail[d.K] = d.V
	}

	return newBerrDetailMap(errorType, errorMessage, errorDetail)
}

func newBerrDetailMap(errorType berrconst.BerrType, errorMessage string, errorDetail map[string]any) *berr {
	return &berr{
		errorType:    errorType,
		errorMessage: errorMessage,
		errorDetail:  errorDetail,
	}
}

type berr struct {
	errorType    berrconst.BerrType
	errorMessage string
	errorDetail  map[string]any
	nextError    error
}

func (e *berr) Error() string {
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

func (e *berr) Unwrap() error {
	return e.nextError
}

// func (e *Berr) Is(target error) bool {
// 	// TODO: implement
// 	return false
// }

func (e *berr) ErrorType() berrconst.BerrType {
	return e.errorType
}

func (e *berr) ErrorMessage() string {
	return e.errorMessage
}

func (e *berr) ErrorDetail() map[string]any {
	if e.errorDetail != nil && len(e.errorDetail) > 0 {
		detailCopy := make(map[string]any)
		for k, v := range e.errorDetail {
			detailCopy[k] = v
		}

		return detailCopy
	}

	return nil
}

func (e *berr) Map() map[string]any {
	return map[string]any{
		"error_type": e.ErrorType().String(),
		"message":    e.ErrorMessage(),
		"detail":     e.ErrorDetail(),
	}
}
