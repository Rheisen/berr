package berr

// ErrorType denotes common types of errors and additionally provides standard string and HTTP type mappings.
type ErrorType int

const (
	UndefinedErrorType ErrorType = iota
	InvalidErrorType
	ApplicationErrorType
	AuthenticationErrorType
	AuthorizationErrorType
	NotFoundErrorType
	ValueMissingErrorType
	ValueInvalidErrorType
)

var errorTypeStringMap = map[ErrorType]string{
	UndefinedErrorType:      "undefined",
	InvalidErrorType:        "invalid",
	ApplicationErrorType:    "application",
	AuthenticationErrorType: "authentication",
	AuthorizationErrorType:  "authorization",
	NotFoundErrorType:       "not_found",
	ValueInvalidErrorType:   "value_invalid",
	ValueMissingErrorType:   "value_missing",
}

var errorTypeHTTPCodeMap = map[ErrorType]int{
	UndefinedErrorType:      500,
	InvalidErrorType:        500,
	ApplicationErrorType:    500,
	AuthenticationErrorType: 401,
	AuthorizationErrorType:  403,
	NotFoundErrorType:       404,
	ValueInvalidErrorType:   422,
	ValueMissingErrorType:   422,
}

func (e ErrorType) String() string {
	errorTypeStr, found := errorTypeStringMap[e]
	if !found {
		return errorTypeStringMap[InvalidErrorType]
	}

	return errorTypeStr
}

func (e ErrorType) HTTPCode() int {
	errorTypeCode, found := errorTypeHTTPCodeMap[e]
	if !found {
		return errorTypeHTTPCodeMap[InvalidErrorType]
	}

	return errorTypeCode
}
