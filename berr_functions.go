package berr

import (
	"github.com/rheisen/berr/berrconst"
	"github.com/rheisen/berr/internal/berrmodel"
)

func Application(message string, detail map[string]interface{}) *berrmodel.Berr {
	return berrmodel.NewBerr(berrconst.ApplicationErrorType, message, detail)
}

func ValueInvalid(message string, detail map[string]interface{}) *berrmodel.Berr {
	return berrmodel.NewBerr(berrconst.ValueInvalidErrorType, message, detail)
}

func ValueMissing(message string, detail map[string]interface{}) *berrmodel.Berr {
	return berrmodel.NewBerr(berrconst.ValueMissingErrorType, message, detail)
}

func Authorization(message string, detail map[string]interface{}) *berrmodel.Berr {
	return berrmodel.NewBerr(berrconst.AuthorizationErrorType, message, detail)
}

func Authentication(message string, detail map[string]interface{}) *berrmodel.Berr {
	return berrmodel.NewBerr(berrconst.AuthenticationErrorType, message, detail)
}

func NotFound(message string, detail map[string]interface{}) *berrmodel.Berr {
	return berrmodel.NewBerr(berrconst.NotFoundErrorType, message, detail)
}
