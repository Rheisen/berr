package berrconst

// BerrType denotes common types of errors and additionally provides standard string and HTTP type mappings.
type BerrType int

const (
	UndefinedBerrType BerrType = iota
	InvalidBerrType
	ApplicationErrorType
	AuthenticationErrorType
	AuthorizationErrorType
	NotFoundErrorType
	ValueMissingErrorType
	ValueInvalidErrorType
)

var berrTypeToStringMap = map[BerrType]string{
	UndefinedBerrType:       "undefined",
	InvalidBerrType:         "invalid",
	ApplicationErrorType:    "application",
	AuthenticationErrorType: "authentication",
	AuthorizationErrorType:  "authorization",
	NotFoundErrorType:       "not_found",
	ValueInvalidErrorType:   "value_invalid",
	ValueMissingErrorType:   "value_missing",
}

// NOTE: Undefined and Invalid Berr Types return 500 HTTP Status Codes, as the 500 HTTP Status Code most closely
// matches the situation of an Undefined or Invalid Berr being returned from an applicaiton.
var berrHTTPCodeMap = map[BerrType]int{
	UndefinedBerrType:       500,
	InvalidBerrType:         500,
	ApplicationErrorType:    500,
	AuthenticationErrorType: 401,
	AuthorizationErrorType:  403,
	NotFoundErrorType:       404,
	ValueInvalidErrorType:   422,
	ValueMissingErrorType:   422,
}

func (e BerrType) String() string {
	errorTypeStr, found := berrTypeToStringMap[e]
	if !found {
		return berrTypeToStringMap[InvalidBerrType]
	}

	return errorTypeStr
}

func (e BerrType) HTTPCode() int {
	errorTypeCode, found := berrHTTPCodeMap[e]
	if !found {
		return berrHTTPCodeMap[InvalidBerrType]
	}

	return errorTypeCode
}
