package berr

import (
	"github.com/rheisen/berr/berrconst"
)

func Application(message string, details ...D) Error {
	return newBerr(berrconst.ApplicationErrorType, message, details...)
}

func ValueInvalid(message string, details ...D) Error {
	return newBerr(berrconst.ValueInvalidErrorType, message, details...)
}

func ValueMissing(message string, details ...D) Error {
	return newBerr(berrconst.ValueMissingErrorType, message, details...)
}

func Authorization(message string, details ...D) Error {
	return newBerr(berrconst.AuthorizationErrorType, message, details...)
}

func Authentication(message string, details ...D) Error {
	return newBerr(berrconst.AuthenticationErrorType, message, details...)
}

func NotFound(message string, details ...D) Error {
	return newBerr(berrconst.NotFoundErrorType, message, details...)
}
