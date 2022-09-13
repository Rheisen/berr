package berr

import (
	"github.com/rheisen/berr/berrconst"
)

func Application(message string, detail map[string]interface{}) *Berr {
	return newBerr(berrconst.ApplicationErrorType, message, detail)
}

func ValueInvalid(message string, detail map[string]interface{}) *Berr {
	return newBerr(berrconst.ValueInvalidErrorType, message, detail)
}

func ValueMissing(message string, detail map[string]interface{}) *Berr {
	return newBerr(berrconst.ValueMissingErrorType, message, detail)
}

func Authorization(message string, detail map[string]interface{}) *Berr {
	return newBerr(berrconst.AuthorizationErrorType, message, detail)
}

func Authentication(message string, detail map[string]interface{}) *Berr {
	return newBerr(berrconst.AuthenticationErrorType, message, detail)
}

func NotFound(message string, detail map[string]interface{}) *Berr {
	return newBerr(berrconst.NotFoundErrorType, message, detail)
}
