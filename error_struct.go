package berr

import (
	"fmt"
	"strings"
)

func newBerr(errorType ErrorType, errorMessage string, attachments ...*Attachment) *Error {
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
		case d.Sensitive():
			errorMetadata[d.Key()] = d.Value()
		default:
			errorDetail[d.Key()] = d.Value()
		}
	}

	return newBerrWithAttachments(errorType, errorMessage, errorDetail, errorMetadata, next)
}

func newBerrWithAttachments(
	errorType ErrorType, errorMessage string, errorDetail, errorMetadata map[string]any, next error,
) *Error {
	return &Error{
		ErrType:       errorType,
		ErrTypeString: errorType.String(),
		ErrMessage:    errorMessage,
		ErrDetails:    errorDetail,
		ErrMetadata:   errorMetadata,
		nextError:     next,
	}
}

type Errors []*Error

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

type Error struct {
	ErrDetails    map[string]any `json:"details"`
	ErrMetadata   map[string]any `json:"-"`
	nextError     error
	ErrTypeString string    `json:"error_type"`
	ErrMessage    string    `json:"message"`
	ErrType       ErrorType `json:"-"`
}

func (e Error) String() string {
	return fmt.Sprintf("[%s_error] %s", e.ErrTypeString, e.ErrMessage)
}

func (e Error) Error() string {
	if e.nextError != nil {
		return fmt.Sprintf("%s: %s", e.String(), e.nextError.Error())
	}

	return e.String()
}

func (e Error) Unwrap() error {
	return e.nextError
}

func (e Error) Type() ErrorType {
	return e.ErrType
}

func (e Error) HTTPCode() int {
	return e.ErrType.HTTPCode()
}

func (e Error) Message() string {
	return e.ErrMessage
}

func (e Error) Details() map[string]any {
	if e.ErrDetails != nil && len(e.ErrDetails) > 0 {
		detailCopy := make(map[string]any)
		for k, v := range e.ErrDetails {
			detailCopy[k] = v
		}

		return detailCopy
	}

	return nil
}

func (e Error) Metadata() map[string]any {
	if e.ErrMetadata != nil && len(e.ErrMetadata) > 0 {
		metadataCopy := make(map[string]any)
		for k, v := range e.ErrMetadata {
			metadataCopy[k] = v
		}

		return metadataCopy
	}

	return nil
}

func (e Error) Map() map[string]any {
	return map[string]any{
		"error_type": e.Type().String(),
		"message":    e.Message(),
		"details":    e.Details(),
	}
}

func (e Error) FullMap() map[string]any {
	return map[string]any{
		"error_type": e.Type().String(),
		"message":    e.Message(),
		"details":    e.Details(),
		"metadata":   e.Metadata(),
	}
}
