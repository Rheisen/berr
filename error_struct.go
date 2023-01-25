package berr

import (
	"fmt"

	"github.com/rheisen/berr/berrconst"
)

func newBerr(errorType berrconst.BerrType, errorMessage string, details ...Detail) *berr {
	errorDetail := make(map[string]any, len(details))
	var next error
	for _, d := range details {
		if d.Type() == "berr_error_detail" && next == nil {
			next, _ = d.Value().(error)
		} else if d.Type() == "berr_error_detail" {
			next = fmt.Errorf("%w: %w", next, d.Value().(error))
		} else {
			errorDetail[d.Key()] = d.Value()
		}
	}

	return newBerrDetailMap(errorType, errorMessage, errorDetail, next)
}

func newBerrDetailMap(errorType berrconst.BerrType, errorMessage string, errorDetail map[string]any, next error) *berr {
	return &berr{
		Type:       errorType,
		TypeString: errorType.String(),
		Message:    errorMessage,
		Detail:     errorDetail,
		nextError:  next,
	}
}

type berr struct {
	Type       berrconst.BerrType `json:"-"`
	TypeString string             `json:"error_type"`
	Message    string             `json:"message"`
	Detail     map[string]any     `json:"detail"`
	nextError  error
}

func (e *berr) Error() string {
	if e.Detail != nil && len(e.Detail) > 0 {
		return fmt.Sprintf(
			"%s: %s (%v)",
			e.Type,
			e.Message,
			e.Detail,
		)
	}

	return fmt.Sprintf(
		"%s: %s",
		e.Type,
		e.Message,
	)
}

func (e *berr) Unwrap() error {
	return e.nextError
}

// func (e *berr) Is(target error) bool {
// 	if target == e {
// 		return true
// 	}

// 	return false
// }

func (e *berr) ErrorType() berrconst.BerrType {
	return e.Type
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
