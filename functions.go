package berr

func Application(message string, details ...Detail) Error {
	return newBerr(ApplicationErrorType, message, details...)
}

func ValueInvalid(message string, details ...Detail) Error {
	return newBerr(ValueInvalidErrorType, message, details...)
}

func ValueMissing(message string, details ...Detail) Error {
	return newBerr(ValueMissingErrorType, message, details...)
}

func Authorization(message string, details ...Detail) Error {
	return newBerr(AuthorizationErrorType, message, details...)
}

func Authentication(message string, details ...Detail) Error {
	return newBerr(AuthenticationErrorType, message, details...)
}

func NotFound(message string, details ...Detail) Error {
	return newBerr(NotFoundErrorType, message, details...)
}

func New(errType ErrorType, message string, details ...Detail) Error {
	return newBerr(errType, message, details...)
}

func D(key string, value any) Detail {
	return detail{key: key, value: value}
}

func E(value error) Detail {
	return errorDetail{key: "error", value: value}
}
