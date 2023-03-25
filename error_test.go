package berr_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rheisen/berr"
)

func TestApplicationError(t *testing.T) {
	const errorMessage = "unexpected problem unmarshalling struct"

	err := berr.Application(errorMessage)
	if err == nil {
		t.Fatalf("unexpected nil berr.Error")
	}

	if err.Error() != err.String() {
		t.Errorf("expected err.Error() '%s', to equal err.String() '%s'", err.Error(), err.String())
	}

	if !strings.Contains(err.Error(), errorMessage) {
		t.Errorf("unexpected err.Error() value: %s", err.Error())
	}

	if !strings.Contains(err.Error(), fmt.Sprintf("[%s error]", berr.ApplicationErrorType.String())) {
		t.Errorf("unexpected err.Error() value: %s", err.Error())
	}

	if len(err.Details()) > 0 {
		t.Errorf("unexpected length of err.Details(): '%d'", len(err.Details()))
	}

	if val, found := err.Map()["error_type"]; !found {
		t.Errorf("expected err.Map() to contain 'error_type' key")
	} else if val != berr.ApplicationErrorType.String() {
		t.Errorf("unexpected error_type value in err.Map(): '%s'", val)
	}

	if _, found := err.Map()["details"]; !found {
		t.Errorf("expected err.Map() to contain 'details' key")
	}

	if _, found := err.Map()["message"]; !found {
		t.Errorf("expected err.Map() to contain 'message' key")
	}
}

func TestValueInvalidError(t *testing.T) {
}

func TestValueMissingError(t *testing.T) {
}

func TestNotFoundError(t *testing.T) {
}

func TestAuthorizationError(t *testing.T) {
}

func TestAuthenticationError(t *testing.T) {
}

func TestApplicationErrorNoDetails(t *testing.T) {
	errorMessage := "message"
	err := berr.Application(errorMessage)

	if err.Type() != berr.ApplicationErrorType {
		t.Errorf(
			"unexpected application berr error_type: expected '%s', found '%s'",
			berr.ApplicationErrorType.String(),
			err.Type().String(),
		)
	}

	if err.Message() != errorMessage {
		t.Errorf(
			"unexpected application berr error_message: expected '%s', found '%s'",
			errorMessage,
			err.Message(),
		)
	}

	expectedError := fmt.Sprintf("[%s error] %s", berr.ApplicationErrorType.String(), errorMessage)
	if err.Error() != expectedError {
		t.Errorf(
			"unexpected application berr error: expected '%s', found '%s'",
			expectedError,
			err.Error(),
		)
	}

	if err.Details() != nil {
		t.Errorf(
			"unexpected application berr error_detail: expected nil, found '%v'",
			err.Details(),
		)
	}
}

func TestApplicationErrorWithDetails(t *testing.T) {
	errorMessage := "message"
	errDetailA := berr.D("some", 2)

	err := berr.Application(errorMessage, errDetailA)

	if err.Message() != errorMessage {
		t.Errorf(
			"unexpected application berr error_message: expected '%s', found '%s'",
			errorMessage,
			err.Message(),
		)
	}
}

func TestApplicationErrorWithErrorDetail(t *testing.T) {
	errorMessage := "problem creating client"

	err := berr.Application(errorMessage, berr.E(fmt.Errorf("problem pinging host")))

	t.Log(err.String())
	t.Log(err)
}
