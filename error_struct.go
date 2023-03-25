package berr

import (
	"fmt"
)

func newBerr(errorType ErrorType, errorMessage string, details ...Detail) *berr {
	errorDetail := make(map[string]any, len(details))
	var next error
	for _, d := range details {
		if d.Type() == "berr_error_detail" && next == nil {
			next, _ = d.Value().(error)
		} else if d.Type() == "berr_error_detail" {
			next = fmt.Errorf("%s: %w", next, d.Value().(error))
		} else {
			errorDetail[d.Key()] = d.Value()
		}
	}

	return newBerrDetailMap(errorType, errorMessage, errorDetail, next)
}

func newBerrDetailMap(errorType ErrorType, errorMessage string, errorDetail map[string]any, next error) *berr {
	return &berr{
		ErrType:       errorType,
		ErrTypeString: errorType.String(),
		ErrMessage:    errorMessage,
		ErrDetails:    errorDetail,
		nextError:     next,
	}
}

type berr struct {
	ErrType       ErrorType      `json:"-"`
	ErrTypeString string         `json:"error_type"`
	ErrMessage    string         `json:"message"`
	ErrDetails    map[string]any `json:"details"`
	nextError     error
}

func (e *berr) String() string {
	return fmt.Sprintf("[%s error] %s", e.ErrTypeString, e.ErrMessage)
}

func (e *berr) Error() string {
	if e.nextError != nil {
		return fmt.Sprintf("%s: %s", e.String(), e.nextError.Error())
	}

	return e.String()
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

func (e *berr) Type() ErrorType {
	return e.ErrType
}

func (e *berr) Message() string {
	return e.ErrMessage
}

func (e *berr) Details() map[string]any {
	if e.ErrDetails != nil && len(e.ErrDetails) > 0 {
		detailCopy := make(map[string]any)
		for k, v := range e.ErrDetails {
			detailCopy[k] = v
		}

		return detailCopy
	}

	return nil
}

func (e *berr) Map() map[string]any {
	return map[string]any{
		"error_type": e.Type().String(),
		"message":    e.Message(),
		"details":    e.Details(),
	}
}
