package berr_test

import (
	"strings"
	"testing"

	"github.com/rheisen/berr"
)

func TestErrorStructThroughInterface(t *testing.T) {
	const errorMessage = "unexpected problem unmarshalling struct"
	const nestedErrorMessage = "invalid struct tag"

	const metadataAttachKey = "metadata_key"
	const metadataAttachValue = "metadata_val"

	const detailAttachKey = "detail_key"
	const detailAttachValue = "detail_val"

	metadataAttachment := berr.M(metadataAttachKey, metadataAttachValue)
	detailAttachment := berr.D(detailAttachKey, detailAttachValue)
	errorAttachment := berr.E(berr.ValueInvalid(nestedErrorMessage))

	err := berr.Application(errorMessage, metadataAttachment, detailAttachment, errorAttachment)
	if err == nil {
		t.Fatalf("unexpected nil berr.Error\n")
	}

	// -- Test err.Message() value --

	if err.Message() != errorMessage {
		t.Errorf("err.Message() = '%s', expected '%s'\n", err.Message(), errorMessage)
	}

	// -- Test err.String() value --

	if !strings.Contains(err.String(), err.Message()) {
		t.Errorf("err.String() does not contain err.message(), found: '%s'\n", err.String())
	}

	if !strings.Contains(err.String(), "[application_error]") {
		t.Errorf("err.String() missing error type identifier, found: '%s'\n", err.String())
	}

	if strings.Contains(err.String(), nestedErrorMessage) {
		t.Errorf("err.String() contains nested error message, found: '%s'\n", err.String())
	}

	// -- Test err.Error() value --

	if !strings.Contains(err.Error(), err.String()) {
		t.Errorf("err.Error() does not contain err.String(), found: '%s'\n", err.Error())
	}

	if !strings.Contains(err.Error(), nestedErrorMessage) {
		t.Errorf("err.Error() does not contain nested error, found: '%s'\n", err.Error())
	}

	// -- Test err.Details() value --
	details := err.Details()

	if len(details) != 1 {
		t.Fatalf("len(err.Details()) does not match expected (1), found: %d\n", len(details))
	}

	detailVal, found := details[detailAttachKey]
	if !found {
		t.Fatalf("detail not found with key '%s'\n", detailAttachKey)
	}

	if detailVal != detailAttachValue {
		t.Errorf(
			"unexpected detail value for key '%s' (expected '%s') found: %s\n",
			detailAttachKey, detailAttachValue, detailVal,
		)
	}

	// -- Test err.Metadata() value --
	metadata := err.Metadata()

	if len(metadata) != 1 {
		t.Fatalf("len(err.Metadata()) does not match expected (1), found: %d\n", len(metadata))
	}

	metadataVal, found := metadata[metadataAttachKey]
	if !found {
		t.Fatalf("metadata not found with key '%s'\n", metadataAttachKey)
	}

	if metadataVal != metadataAttachValue {
		t.Errorf(
			"unexpected metadata value for key '%s' (expected '%s') found: %s\n",
			metadataAttachKey, metadataAttachValue, metadataVal,
		)
	}

	// -- Test err.Unwrap() value --
	wrappedErr := err.Unwrap()

	if wrappedErr == nil {
		t.Fatalf("unexpected nil value found for err.Unwrap()\n")
	}

	if !strings.Contains(wrappedErr.Error(), nestedErrorMessage) {
		t.Errorf(
			"unexpected error message found in unwrapped err (expected to contain '%s'), found: %s\n",
			nestedErrorMessage, wrappedErr.Error(),
		)
	}

	// -- Test err.Map() value --
	errMap := err.Map()

	if len(errMap) != 3 {
		t.Fatalf("len(err.Map()) does not match expected (3), found: %d\n", len(errMap))
	}

	if _, found := errMap["details"]; !found {
		t.Errorf("expected 'details' key in err.Map()\n")
	}

	if _, found := errMap["error_type"]; !found {
		t.Errorf("expected 'error_type' key in err.Map()\n")
	}

	if _, found := errMap["message"]; !found {
		t.Errorf("expected 'message' key in err.Map()\n")
	}

	// -- Test err.FullMap() value --
	fullErrMap := err.FullMap()

	if len(fullErrMap) != 4 {
		t.Fatalf("len(err.FullMap()) does not match expected (4), found: %d\n", len(fullErrMap))
	}

	if _, found := fullErrMap["details"]; !found {
		t.Errorf("expected 'details' key in err.FullMap()\n")
	}

	if _, found := fullErrMap["metadata"]; !found {
		t.Errorf("expected 'metadata' key in err.FullMap()\n")
	}

	if _, found := fullErrMap["error_type"]; !found {
		t.Errorf("expected 'error_type' key in err.FullMap()\n")
	}

	if _, found := fullErrMap["message"]; !found {
		t.Errorf("expected 'message' key in err.FullMap()\n")
	}

	// -- Test err.HTTPCode() value --
	if err.HTTPCode() != berr.ApplicationErrorType.HTTPCode() {
		t.Errorf(
			"unexpected err.HTTPCode() value (expected '%d'), found: %d\n",
			berr.ApplicationErrorType.HTTPCode(), err.HTTPCode(),
		)
	}

	// -- Test err.Type() value --
	if err.Type() != berr.ApplicationErrorType {
		t.Errorf("unexpected err.Type() value (expected '%s'), found: %s\n", berr.ApplicationErrorType, err.Type())
	}
}

// func TestValueInvalidError(t *testing.T) {
// }

// func TestValueMissingError(t *testing.T) {
// }

// func TestNotFoundError(t *testing.T) {
// }

// func TestAuthorizationError(t *testing.T) {
// }

// func TestAuthenticationError(t *testing.T) {
// }

// func TestApplicationErrorNoDetails(t *testing.T) {
// 	errorMessage := "message"
// 	err := berr.Application(errorMessage)

// 	if err.Type() != berr.ApplicationErrorType {
// 		t.Errorf(
// 			"unexpected application berr error_type: expected '%s', found '%s'",
// 			berr.ApplicationErrorType.String(),
// 			err.Type().String(),
// 		)
// 	}

// 	if err.Message() != errorMessage {
// 		t.Errorf(
// 			"unexpected application berr error_message: expected '%s', found '%s'",
// 			errorMessage,
// 			err.Message(),
// 		)
// 	}

// 	expectedError := fmt.Sprintf("[%s error] %s", berr.ApplicationErrorType.String(), errorMessage)
// 	if err.Error() != expectedError {
// 		t.Errorf(
// 			"unexpected application berr error: expected '%s', found '%s'",
// 			expectedError,
// 			err.Error(),
// 		)
// 	}

// 	if err.Details() != nil {
// 		t.Errorf(
// 			"unexpected application berr error_detail: expected nil, found '%v'",
// 			err.Details(),
// 		)
// 	}
// }

// func TestApplicationErrorWithDetails(t *testing.T) {
// 	errorMessage := "message"
// 	errDetailA := berr.D("some", 2)

// 	err := berr.Application(errorMessage, errDetailA)

// 	if err.Message() != errorMessage {
// 		t.Errorf(
// 			"unexpected application berr error_message: expected '%s', found '%s'",
// 			errorMessage,
// 			err.Message(),
// 		)
// 	}
// }

// func TestApplicationErrorWithErrorDetail(t *testing.T) {
// 	errorMessage := "problem creating client"

// 	err := berr.Application(errorMessage, berr.E(fmt.Errorf("problem pinging host")))

// 	t.Log(err.String())
// 	t.Log(err)
// }
