package berr

func FromError(err error) (Error, bool) {
	berrError, okay := err.(Error)

	return berrError, okay
}

func New(errorType ErrorType, message string, attachments ...Attachment) Error {
	return newBerr(errorType, message, attachments...)
}

func Application(message string, attachments ...Attachment) Error {
	return New(ApplicationErrorType, message, attachments...)
}

func ValueInvalid(message string, attachments ...Attachment) Error {
	return New(ValueInvalidErrorType, message, attachments...)
}

func ValueMissing(message string, attachments ...Attachment) Error {
	return New(ValueMissingErrorType, message, attachments...)
}

func Authorization(message string, attachments ...Attachment) Error {
	return New(AuthorizationErrorType, message, attachments...)
}

func Authentication(message string, attachments ...Attachment) Error {
	return New(AuthenticationErrorType, message, attachments...)
}

func NotFound(message string, attachments ...Attachment) Error {
	return New(NotFoundErrorType, message, attachments...)
}

func Unimplemented(message string, attachments ...Attachment) Error {
	return New(UnimplementedErrorType, message, attachments...)
}

func Timeout(message string, attachments ...Attachment) Error {
	return New(TimeoutErrorType, message, attachments...)
}

func Detail(key string, value any) Attachment {
	return detailAttachment{key: key, value: value}
}

func D(key string, value any) Attachment {
	return Detail(key, value)
}

func Metadata(key string, value any) Attachment {
	return metadataAttachment{key: key, value: value}
}

func M(key string, value any) Attachment {
	return Metadata(key, value)
}

func Err(value error) Attachment {
	return errorAttachment{key: "error", value: value}
}

func E(value error) Attachment {
	return Err(value)
}
