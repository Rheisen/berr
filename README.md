# `berr`: better errors for go

`berr` is an errors package that provides simple functions for creating more descriptive errors.

```
go get github.com/rheisen/berr
```

### Why `berr`

When handling errors in Go you typically encounter either the standard library error (a string for all intents and
purposes), or a custom error struct tailored for a specific use case. The `berr` package sits in between, providing
general purpose error structures that provide more descriptive errors with minimal configuration, with a host of options
for creating really powerful errors.

### The `berr` Way

1. Errors need to serve the needs of different audiences, audiences with often conflicting interests (E.g. developers
vs. end-users of APIs and Applications). When creating a `berr.Error`, provide a message that is safe and valuable to
end-users, and add `berr.Attachment` structures to provide additional context relevant for those different audiences.
By default, berr provides `detail`, `metadata`, and `error` attachments. The `detail` attachment is not sensitive, and
intended for end-users, while `metadata` and `error` attachments are sensitive (intended for developers).

2. Errors should be sensitive by default, and not reveal or leak sensitive information / sensitive attachments without
an explicit call to do so. `berr.Error.String()` will output the error type and the message it was initially suplied
with, while `berr.Error.Error()` will output the `berr.Error.String()` in addition to the output of `.Error()` on the
next error (if one exists). A JSON rendering of a `berr.Error` model will only expose the message, error type, and
attachments that are not sensitive (e.g. `detail` attachments).

3. Additional error information should be readily available. If you would like to access metadata, it is available on
the `berr.Error` interface with the `.Metadata()` method, and will also be output with the `.FullMap()` method.

### Creating `berr.Error` Models

* `berr.New(errorType berr.ErrorType, message string, details ...berr.Detail) berr.Error`
* `berr.Application(message string, details ...berr.Detail) berr.Error`
* `berr.Authentication(message string, details ...berr.Detail) berr.Error`
* `berr.Authorization(message string, details ...berr.Detail) berr.Error`
* `berr.NotFound(message string, details ...berr.Detail) berr.Error`
* `berr.ValueInvalid(message string, details ...berr.Detail) berr.Error`
* `berr.ValueMissing(message string, details ...berr.Detail) berr.Error`
* `berr.Unimplemented(message string, details ...berr.Detail) berr.Error`
* `berr.Timeout(message string, details ...berr.Detail) berr.Error`

### Error Types

* `berr.ApplicationErrorType`
* `berr.AuthenticationErrorType`
* `berr.AuthorizationErrorType`
* `berr.NotFoundErrorType`
* `berr.ValueInvalidErrorType`
* `berr.ValueMissingErrorType`
* `berr.UnimplementedErrorType`
* `berr.TimeoutErrorType`

### Getting Information from a `berr.Error`

* `Error.Type() ErrorType`
* `Error.Message() string`
* `Error.Details() map[string]any`
* `Error.Metadata() map[string]any`
* `Error.Code() int`
* `Error.Map() map[string]any`
* `Error.FullMap() map[string]any`
* `Error.Unwrap()`

### Examples

```go
```
