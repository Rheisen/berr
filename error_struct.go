package berr

import (
	"fmt"
)

func newBerr(errorType ErrorType, errorMessage string, attachments ...Attachment) *berr {
	errorDetail := make(map[string]any)
	errorMetadata := make(map[string]any)

	var next error

	for _, d := range attachments {
		switch {
		case d.Type() == AttachmentErrorType && next == nil:
			next, _ = d.Value().(error)
		case d.Type() == AttachmentErrorType:
			next = fmt.Errorf("%s: %w", next, d.Value().(error))
		case d.Type() == AttachmentMetadataType:
			errorMetadata[d.Key()] = d.Value()
		default:
			errorDetail[d.Key()] = d.Value()
		}
	}

	return newBerrWithAttachments(errorType, errorMessage, errorDetail, errorMetadata, next)
}

func newBerrWithAttachments(
	errorType ErrorType, errorMessage string, errorDetail, errorMetadata map[string]any, next error,
) *berr {
	return &berr{
		ErrType:       errorType,
		ErrTypeString: errorType.String(),
		ErrMessage:    errorMessage,
		ErrDetails:    errorDetail,
		ErrMetadata:   errorMetadata,
		nextError:     next,
	}
}

type berr struct {
	ErrDetails    map[string]any `json:"details"`
	ErrMetadata   map[string]any `json:"-"`
	nextError     error
	ErrTypeString string    `json:"error_type"`
	ErrMessage    string    `json:"message"`
	ErrType       ErrorType `json:"-"`
}

func (e *berr) String() string {
	return fmt.Sprintf("[%s_error] %s", e.ErrTypeString, e.ErrMessage)
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

func (e *berr) Type() ErrorType {
	return e.ErrType
}

func (e *berr) HTTPCode() int {
	return e.ErrType.HTTPCode()
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

func (e *berr) Metadata() map[string]any {
	if e.ErrMetadata != nil && len(e.ErrMetadata) > 0 {
		metadataCopy := make(map[string]any)
		for k, v := range e.ErrMetadata {
			metadataCopy[k] = v
		}

		return metadataCopy
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

func (e *berr) FullMap() map[string]any {
	return map[string]any{
		"error_type": e.Type().String(),
		"message":    e.Message(),
		"details":    e.Details(),
		"metadata":   e.Metadata(),
	}
}
