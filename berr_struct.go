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
		BerrType: errorType,
		Message:  errorMessage,
		Detail:   errorDetail,
	}
}

type berr struct {
	BerrType  berrconst.BerrType `json:"error_type"`
	Message   string             `json:"message"`
	Detail    map[string]any     `json:"detail"`
	nextError error
}

func (e *berr) Error() string {
	if e.Detail != nil && len(e.Detail) > 0 {
		return fmt.Sprintf(
			"%s: %s (%v)",
			e.BerrType,
			e.Message,
			e.Detail,
		)
	}

	return fmt.Sprintf(
		"%s: %s",
		e.BerrType,
		e.Message,
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
	return e.BerrType
}

func (e *berr) ErrorMessage() string {
	return e.Message
}

func (e *berr) ErrorDetail() map[string]any {
	if e.Detail != nil && len(e.Detail) > 0 {
		detailCopy := make(map[string]any)
		for k, v := range e.Detail {
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
