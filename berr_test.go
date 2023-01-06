package berr_test

import (
	"fmt"
	"testing"

	"github.com/rheisen/berr"
	"github.com/rheisen/berr/berrconst"
)

func TestApplicationBerrNoDetails(t *testing.T) {
	errorMessage := "message"
	err := berr.Application(errorMessage)

	if err.ErrorType() != berrconst.ApplicationErrorType {
		t.Errorf(
			"unexpected application berr error_type: expected '%s', found '%s'",
			berrconst.ApplicationErrorType.String(),
			err.ErrorType().String(),
		)
	}

	if err.ErrorMessage() != errorMessage {
		t.Errorf(
			"unexpected application berr error_message: expected '%s', found '%s'",
			errorMessage,
			err.ErrorMessage(),
		)
	}

	expectedError := fmt.Sprintf("%s: %s", berrconst.ApplicationErrorType.String(), errorMessage)
	if err.Error() != expectedError {
		t.Errorf(
			"unexpected application berr error: expected '%s', found '%s'",
			expectedError,
			err.Error(),
		)
	}

	if err.ErrorDetail() != nil {
		t.Errorf(
			"unexpected application berr error_detail: expected nil, found '%v'",
			err.ErrorDetail(),
		)
	}
}

func TestApplicationBerrWithDetails(t *testing.T) {
	errorMessage := "message"
	errDetailA := berr.D{K: "some", V: 2}

	err := berr.Application(errorMessage, errDetailA)

	if err.ErrorMessage() != errorMessage {
		t.Errorf(
			"unexpected application berr error_message: expected '%s', found '%s'",
			errorMessage,
			err.ErrorMessage(),
		)
	}
}
